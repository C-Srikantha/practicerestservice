package databasecon

import (
	"github.com/go-pg/pg"
)

func Setup() *pg.DB {
	dbdetails := pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     ":8080",
		Database: "moviedatabase",
	}
	con := pg.Connect(&dbdetails)

	return con
}
