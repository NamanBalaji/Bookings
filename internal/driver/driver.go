package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//  DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDbConn = 10
	maxIdleDbConn = 5
	maxDbLifetime = time.Minute * 5
)

// ConnectSQL creates database pool for postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil

}

// testDB tries to ping the database
func testDB(d *sql.DB) error {

	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
