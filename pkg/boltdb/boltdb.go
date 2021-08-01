package boltdb

import (
	"errors"
	"fmt"
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

// UpdateOptions are options for using the Update method
type UpdateOptions struct {
	Bucket string
	Key    string
	Value  string
}

// Update inserts a key value into our db
// boltdb.Update provides a consistent view of the db as part of the transaction
func (c *Client) Update(options UpdateOptions) error {
	key := []byte(options.Key)
	value := []byte(options.Value)

	err := c.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(options.Bucket))
		if err != nil {
			fmt.Println("Failed to create Bucket")
			return err
		}
		err = bucket.Put(key, value)
		if err != nil {
			fmt.Println("Failed to put item")
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Failed to write data")
		return err
	}
	return nil
}

// GetOptions are options for using the Get method
type GetOptions struct {
	Bucket     string
	Key        string
	ValueStore *string
}

// Get returns the value for a given bucket and key
func (c *Client) Get(options GetOptions) error {
	key := []byte(options.Key)
	err := c.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(options.Bucket))
		if bucket == nil {
			return errors.New("could not find Bucket")
		}

		value := bucket.Get(key)
		*options.ValueStore = string(value[:])

		return nil
	})
	if err != nil {
		fmt.Println("Failed to read from Bucket")
		return err
	}
	return nil
}
