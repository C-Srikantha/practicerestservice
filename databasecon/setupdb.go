package databasecon

import (
	"context"

	"github.com/go-pg/pg"
)

func Setup() (*pg.DB, error) {
	dbdetails := pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     ":8080",
		Database: "moviedatabase",
	}
	con := pg.Connect(&dbdetails)
	ctr := context.Background()
	_, err := con.ExecContext(ctr, "SELECT 1")
	return con, err
}
