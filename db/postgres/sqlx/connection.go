package sqlx

import (
	"github.com/jmoiron/sqlx"

	"github.com/inflite-dev/go-utils/db/postgres"
)

// MustOpen wraps sqlx MustOpen, using Config
// Opens connection to a DB and panics on error
func MustOpen(config postgres.Config) *sqlx.DB {
	uri := postgres.BuildConnectionURI(config)
	return sqlx.MustOpen("postgres", uri)
}

// MustConnect wraps sqlx MustConnect, using Config
// Opens connection to a DB, pings, and panics on error
func MustConnect(config postgres.Config) *sqlx.DB {
	uri := postgres.BuildConnectionURI(config)
	return sqlx.MustConnect("postgres", uri)
}
