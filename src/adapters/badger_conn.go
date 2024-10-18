package adapters

import (
	"github.com/dgraph-io/badger"
)

func Connect(path string) (*BadgerUnitGateway, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &BadgerUnitGateway{db: db}, nil
}
