package config

import (
	"database/sql"
	"fmt"

	error "github.com/blohny/helper"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	error.PanicIfError(err)

	err = db.Ping()
	error.PanicIfError(err)

	log.Info().Msg("Connected to database!!")

	return db
}
