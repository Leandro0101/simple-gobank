package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:toor123@localhost:5432/database?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(conn)

	m.Run()
}