package database

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	db *bolt.DB
}

type Response struct {
	Key   string
	Value string
}

//OpenDB set up the database.
func OpenDB() (boltDB *BoltDB, err error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	boltDB = &BoltDB{db}
	if err != nil {
		log.Fatal(err)
	}
	//defer boltDB.db.Close()
	return boltDB, err
}

func InitDB(boltDB *BoltDB) {
	boltDB.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func WriteSomethingToDB(boltDB *BoltDB, key string, value string) *Response {
	status := boltDB.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
	if status == nil {
		return &Response{key, value}
	} else {

		return &Response{key, "error"}
	}
}

func RetrieveSomethingFromDB(boltDB *BoltDB, key string) *Response {
	var value string
	boltDB.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		value = string(v)
		return nil
	})
	return &Response{key, value}
}
