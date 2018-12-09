package main

import (
	"errors"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func save(bucket string, json string) {

	var world = []byte(bucket)

	db, err := bolt.Open("/tmp/tempIO.db", 0644, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	currentTime := time.Now()

	key := []byte(currentTime.Format("2006-01-02_15:04"))
	value := []byte(json)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

func read(bucketS string, key string) string {
	db, err := bolt.Open("/tmp/tempIO.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	value := ""

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketS))
		if bucket == nil {
			return errors.New("Bucket not found")
		}

		val := bucket.Get([]byte(key))

		value = string(val)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value
}
