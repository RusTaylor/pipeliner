package migration

import (
	"fmt"
	"log"
	"os"
	"pipeliner/database"
)

func Migrate() {
	log.Println("Migrate start")
	db, err := database.GetDb()

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	err = db.Query(`
CREATE TABLE IF NOT EXISTS "user"
(
    id       serial
        constraint user_pk
            primary key,
    name     varchar,
    login    varchar,
    password varchar
)
`)
	if err != nil {
		log.Println("Migrate error: " + fmt.Sprintf("%v", err))
		os.Exit(2)
	}

	log.Println("Migrate end")
	db.Close()
}
