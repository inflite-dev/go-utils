package sql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/inflite-dev/go-utils/db/postgres"
)

var createDB = `CREATE DATABASE %s`
var dropDB = `DROP DATABASE %s`

func CreateDB(t *testing.T, dbName string, superUserConfig postgres.Config) (*sql.DB, error) {
	t.Helper()

	// Connect to Postgres with user that can create and drop DBs
	superUserDB := MustConnect(superUserConfig)
	fmt.Printf("\nCreating DB %s...\n", dbName)

	createDBStatement := fmt.Sprintf(createDB, dbName)
	_, err := superUserDB.Exec(createDBStatement)
	if err != nil {
		return nil, err
	}

	// Done with the Postgres superuser - close connection
	err = superUserDB.Close()
	if err != nil {
		return nil, err
	}

	// Connect to test DB
	pgTestDBConfig := postgres.Config{
		Host:                  superUserConfig.Host,
		Port:                  superUserConfig.Port,
		Username:              superUserConfig.Username,
		Password:              superUserConfig.Password,
		Database:              dbName,
		ApplicationName:       superUserConfig.ApplicationName,
		ConnectTimeoutSeconds: superUserConfig.ConnectTimeoutSeconds,
		SSLMode:               superUserConfig.SSLMode,
	}

	fmt.Printf("\nOpening DB %s...\n", dbName)
	testDB := MustConnect(pgTestDBConfig)

	return testDB, nil
}

func TearDownDB(t *testing.T, db *sql.DB, superUserConfig postgres.Config) error {
	t.Helper()

	// Extract connected DB name
	var dbName string
	row := db.QueryRow(`SELECT current_catalog;`)
	err := row.Scan(&dbName)
	if err != nil {
		return err
	}

	fmt.Printf("\nClosing DB %s...\n", dbName)
	err = db.Close()
	if err != nil {
		return err
	}

	// Connect to Postgres with user that can drop DBs
	superUserDB := MustConnect(superUserConfig)

	fmt.Printf("\nDropping DB %s...\n", dbName)
	dropDBStatement := fmt.Sprintf(dropDB, dbName)
	_, err = superUserDB.Exec(dropDBStatement)
	if err != nil {
		return err
	}

	// Done with the Postgres superuser - close connection
	err = superUserDB.Close()
	if err != nil {
		return err
	}

	return nil
}
