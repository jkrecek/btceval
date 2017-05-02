package entity

import (
	"time"

	"fmt"

	"github.com/jkrecek/sorm-go"
)

type Record struct {
	sorm.AbstractEntity

	ID        uint      `db:"id" primary:"true"`
	CreatedAt time.Time `db:"created_at"`

	EUR_sell float32 `json:"EURsell" db:"eur_sell"`
	EUR_buy  float32 `json:"EURbuy" db:"eur_buy"`
	CZK_sell float32 `json:"CZKsell" db:"czk_sell"`
	CZK_buy  float32 `json:"CZKbuy" db:"czk_buy"`
	USD_sell float32 `json:"USDsell" db:"usd_sell"`
	USD_buy  float32 `json:"USDbuy" db:"usd_buy"`
}

func (r *Record) String() string {
	return fmt.Sprintf("EUR sell: %f, EUR buy: %f, CZK sell: %f, CZK buy: %f, USD sell: %f, USD buy: %f",
		r.EUR_sell,
		r.EUR_buy,
		r.CZK_sell,
		r.CZK_buy,
		r.USD_sell,
		r.USD_buy,
	)
}
