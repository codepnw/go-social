package main

import (
	"log"

	"github.com/codepnw/social/internal/db"
	"github.com/codepnw/social/internal/env"
	"github.com/codepnw/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}