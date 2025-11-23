package main

import (
	"log"
	"test/internal/env"
	"test/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("addr", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.start(mux))
}
