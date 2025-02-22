package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dnguyen316/GoGetIt/server/internal/config"

	"github.com/doug-martin/goqu/v9"
)

func InitializeDB(databaseConfig config.Database) (db *sql.DB, cleanup func(), err error) {
	// Build the DSN (Data Source Name) for MySQL.
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	// Open a connection using the MySQL driver.
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("failed to open database: %+v\n", err)
		return nil, nil, err
	}

	cleanup = func() {
		db.Close()
	}

	return db, cleanup, nil
}

func InitializeSquirrel(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
}
