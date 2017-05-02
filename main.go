package main

import (
	"log"

	"github.com/jkrecek/btceval/app"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
