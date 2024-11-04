package main

import (
	"log"
	"os"

	"github.com/bombaepabo/social/internal/env"
	"github.com/bombaepabo/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:store,
	}
	os.LookupEnv("PATH")
	mux := app.mount()
	log.Fatal(app.run(mux))
}
