package main

import (
	"log"
	"test/internal/db"
	"test/internal/env"
	"test/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("addr", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://social:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 10),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 5),
			maxLifetime:  env.GetString("DB_MAX_LIFETIME", "10m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxLifetime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("Database connection established!")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.start(mux))
}
