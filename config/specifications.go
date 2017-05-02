package config

type cfgType uint

const (
	DO cfgType = iota
	DB_TYPE
	DB_DSN
	XML_PATH
)

var (
	cfg = map[cfgType]*spec{
		DO: {
			"",
			"do",
			"",
		},
		DB_TYPE: {
			"DB_TYPE",
			"db.type",
			"mysql",
		},
		DB_DSN: {
			"DB_DSN",
			"db.dsn",
			"root:root@tcp(localhost:3306)/db_name",
		},
		XML_PATH: {
			"XML_PATH",
			"xml",
			"prices.xml",
		},
	}
)
