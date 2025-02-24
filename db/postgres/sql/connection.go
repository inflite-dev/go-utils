package sql

import (
	"database/sql"

	// Makes postgres driver available to Golang's database/sql package
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	_ "github.com/lib/pq"

	"github.com/inflite-dev/go-utils/db/postgres"
)

// MustOpen mimics sqlx MustOpen, but for sql.DB, using Config
// Opens connection to a DB and panics on error
func MustOpen(config postgres.Config) *sql.DB {
	uri := postgres.BuildConnectionURI(config)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	return db
}

// MustConnect mimics sqlx MustConnect, but for sql.DB, using Config
// Opens connection to a DB, pings, and panics on error
func MustConnect(config postgres.Config) *sql.DB {
	uri := postgres.BuildConnectionURI(config)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		panic(err)
	}
	return db
}
