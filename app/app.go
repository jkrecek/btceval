package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jkrecek/btceval/config"
	"github.com/jkrecek/btceval/database"
	"github.com/jkrecek/btceval/database/entity"
	"github.com/jkrecek/btceval/database/repository"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	instance *core
)

type core struct {
	rm *repository.Manager

	shouldFinish chan struct{}
	didFinish    chan struct{}
}

func Init() (err error) {
	if instance != nil {
		return
	}

	c := new(core)
	db, err := database.InstanceOptional()
	if err != nil {
		return
	}

	c.rm = repository.NewManager(db)
	c.shouldFinish = make(chan struct{})
	c.didFinish = make(chan struct{})

	instance = c
	return
}

func Run() error {
	if instance == nil {
		return errors.New("Instance was not created")
	}

	instance.run()
	return nil
}

func (c *core) run() {
	ticker := time.NewTicker(15 * time.Second) // TODO config
	c.do()
	go func() {
		for {
			select {
			case <-ticker.C:
				c.do()
			case <-c.shouldFinish:
				ticker.Stop()
				c.didFinish <- struct{}{}
				return
			}
		}
	}()

	<-c.didFinish
}

func (c *core) do() {
	rec, err := loadBTCData()
	if err != nil {
		log.Println(err)
	}

	c.rm.Record().WithEntity(&rec).Save()

	saveToXML(rec)

	fmt.Println(rec.String())
}

func loadBTCData() (res entity.Record, err error) {
	resp, err := http.Get("https://my.wbtcb.com/pub/exchange-rate-now/BTC")
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}

	return
}

func saveToXML(record entity.Record) {
	prices := CreatePricesEntity(record)
	bts, err := xml.Marshal(prices)
	if err != nil {
		log.Println(err)
	}

	file, err := os.Create(config.GetValue(config.XML_PATH))
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	_, err = file.WriteString(xml.Header)
	if err != nil {
		log.Println(err)
	}

	_, err = file.Write(bts)
	if err != nil {
		log.Println(err)
	}

	err = file.Sync()
	if err != nil {
		log.Println(err)
	}
}
