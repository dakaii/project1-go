package database

import (
	"log"
	"github.com/boltdb/bolt"
)

//OpenDB set up the database.
func OpenDB() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
