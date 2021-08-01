package ports

type Database interface {
	Open() error
	Close()
	Update(bucket, key, value string) error
	Get(bucket, key string) (string, error)
}
