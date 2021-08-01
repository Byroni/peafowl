package ports

import "github.com/byroni/peafowl/pkg/boltdb"

type Database interface {
	Open() error
	Close()
	Update(options boltdb.UpdateOptions) error
	Get(options boltdb.GetOptions) error
}
