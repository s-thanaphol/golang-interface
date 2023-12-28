package database

type Database interface {
	Get() (int, error)
	Post(int) error
}