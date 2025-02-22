package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dnguyen316/GoGetIt/server/internal/config"

	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql" // Import MySQL goqu dialect
	_ "github.com/go-sql-driver/mysql"               // Import MySQL driver
)

type Database interface {
	Delete(table interface{}) *goqu.DeleteDataset
	Dialect() string
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	From(from ...interface{}) *goqu.SelectDataset
	Insert(table interface{}) *goqu.InsertDataset
	Logger(logger goqu.Logger)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructs(i interface{}, query string, args ...interface{}) error
	ScanStructsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	ScanVal(i interface{}, query string, args ...interface{}) (bool, error)
	ScanValContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanVals(i interface{}, query string, args ...interface{}) error
	ScanValsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	Select(cols ...interface{}) *goqu.SelectDataset
	Trace(op string, sqlString string, args ...interface{})
	Truncate(table ...interface{}) *goqu.TruncateDataset
	Update(table interface{}) *goqu.UpdateDataset
}

func InitializeDB(databaseConfig config.Database) (db *sql.DB, cleanup func(), err error) {
	// Build the DSN (Data Source Name) for MySQL.
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		databaseConfig.AccountName,
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

func InitializeGoquDB(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
}

func ProvideDatabase(config config.Database) (*goqu.Database, func(), error) {
	sqlDB, cleanup, err := InitializeDB(config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize SQL DB: %w", err)
	}

	gdb := InitializeGoquDB(sqlDB)
	return gdb, cleanup, nil
}
