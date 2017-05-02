package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jkrecek/btceval/config"
)

var (
	instance *DB
)

type DB struct {
	*sql.DB
}

func InstanceOptional() (*DB, error) {
	if instance == nil {
		tempInstance, err := establishConnection(
			config.GetValue(config.DB_TYPE),
			config.GetValue(config.DB_DSN),
		)
		if err != nil {
			return nil, err
		}

		instance = &DB{DB: tempInstance}
	}

	return instance, nil
}

func establishConnection(dbType, dsn string) (db *sql.DB, err error) {
	//ud, err := url.Parse(dsn)
	//if err != nil {
	//	return
	//}

	q := make(url.Values)
	q.Set("charset", "utf8")
	q.Set("parseTime", "True")
	q.Set("loc", "Local")
	//ud.RawQuery = q.Encode()

	db, err = sql.Open(dbType, fmt.Sprintf("%s?%s", dsn, q.Encode()))
	if err != nil {
		return
	}

	err = db.Ping()
	return
}
