package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/bitscrow/config"
	"github.com/ngenohkevin/bitscrow/internal/store"
	"log"
	"os"
	"testing"
)

//nolint:unused
var testStore store.Store

// TestMain is the entry point for the tests
func TestMain(m *testing.M) {
	configure, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), configure.DbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = store.NewStore(connPool)
	os.Exit(m.Run())
}
