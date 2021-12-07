package main

import (
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
)

func main() {
	//FIXME: 環境変数で管理できるように
	m, err := migrate.New(
		"file://db/migations",
		"postgres://deploy:password@localhost:5432/deploy?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		if err := m.Steps(1); err != nil {
			log.Fatal(err)
		}
	}
}
