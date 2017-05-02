package app

import (
	"encoding/xml"
	"github.com/jkrecek/btceval/database/entity"
	"time"
)

type Spread struct {
	BuyPrice  float32 `xml:"buy"`
	SellPrice float32 `xml:"sell"`
}

type Prices struct {
	XMLName   xml.Name  `xml:"prices"`
	USD       Spread    `xml:"usd"`
	EUR       Spread    `xml:"eur"`
	CZK       Spread    `xml:"czk"`
	CreatedAt time.Time `xml:"created_at"`
}

func CreatePricesEntity(record entity.Record) Prices {
	return Prices{
		CZK: Spread{
			BuyPrice:  record.CZK_buy,
			SellPrice: record.CZK_sell,
		},
		EUR: Spread{
			BuyPrice:  record.EUR_buy,
			SellPrice: record.EUR_sell,
		},
		USD: Spread{
			BuyPrice:  record.USD_buy,
			SellPrice: record.USD_sell,
		},
		CreatedAt: time.Now(),
	}
}
