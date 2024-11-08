package main

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/bitscrow/config"
	"log"
)

func main() {
	configurations, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connPool, err := pgxpool.New(context.Background(), configurations.DbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	runDBMigrations(configurations.MigrationURL, configurations.DbUrl)

	log.Println("Connected to the database")

	defer connPool.Close()
}

func runDBMigrations(migrationURL string, dbURL string) {
	migration, err := migrate.New(migrationURL, dbURL)
	if err != nil {
		log.Fatal("cannot create migration:", err)
	}
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migration up:", err)
	}
	log.Println("migration successful")
}
