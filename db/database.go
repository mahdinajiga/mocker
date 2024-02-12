package db

//go:generate mockery --name Database
type Database interface {
	Connect() error

	Get(key string) (value string, err error)
	Set(key string, value string) (err error)

	Close() error
}
