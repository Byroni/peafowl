package boltdb

import (
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

type Client struct {
	db *bolt.DB
}

// New returns a new boltdb client
func New() *Client {
	return &Client{}
}

// Open will open a connection
func (c *Client) Open() error {
	db, err := bolt.Open("peafowl.db", 0600, &bolt.Options{
		Timeout:         3 * time.Second,
	})
	if err != nil {
		return err
	}

	c.db = db
	return nil
}

// Close will close the database connection
func (c *Client) Close() {
	c.db.Close()
}

// Update inserts a key value into our db
// boltdb.Update provides a consistent view of the db as part of the transaction
func (c *Client) Update(bucket, key, value string) error {
	k := []byte(key)
	v := []byte(value)

	err := c.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = bucket.Put(k, v)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Get returns the value for a given bucket and key
func (c *Client) Get(bucket, key string) (string, error) {
	var value string
	keyString := []byte(key)
	err := c.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return errors.New("could not find bucket")
		}
		value = string(bucket.Get(keyString))
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, nil
}
