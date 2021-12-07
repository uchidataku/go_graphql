package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations/000001_create_accounts_table",
		"postgres://deploy:password@localhost:5432/deploy?sslmode=disable")
	if err != nil {
		return
	}
	m.Up()
}
