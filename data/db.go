package data

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Database struct {
	Pool *pgxpool.Pool
}

func InitializeDb(connStr string) *Database {
	db := Database{}
	db.initializePool(connStr)
	return &db
}

func (d *Database) initializePool(connStr string) {
	log.Println("Initializing Connection Pool")
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	d.Pool = conn
}

func (d *Database) Close() {
	d.Pool.Close()
}

func (d *Database) CreateTables() {

	_, err := d.Pool.Exec(context.Background(), createUserTable)
	if err != nil {
		log.Fatalf("Creating User table failed: %v\n", err)
	}
	log.Println("Successfully created User table")

	_, err = d.Pool.Exec(context.Background(), createRequestTable)
	if err != nil {
		log.Fatalf("Creating Request table failed: %v\n", err)
	}
	log.Println("Successfully created Requests table")
	log.Println("Tables successfully created")
}
