package heck

import "github.com/gocql/gocql"

func IsCassandraNotFoundError(err error) bool {
	return err == gocql.ErrNotFound
}
