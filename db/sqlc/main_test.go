package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()
	// testDB, err = pgx.Connect(ctx, "host=localhost port=5432 user=postgres password=123123 dbname=simple_bank")
	testDB, err = pgxpool.New(ctx, "host=localhost port=5432 user=postgres password=123123 dbname=simple_bank")

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
