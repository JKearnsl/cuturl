package adapters

import (
	"cuturl/src/domain"
	"github.com/dgraph-io/badger"
)

type BadgerUnitGateway struct {
	db *badger.DB
}

func (b *BadgerUnitGateway) Save(unit *domain.Unit) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(unit.CODE), []byte(unit.URL))
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *BadgerUnitGateway) Get(code string) (*domain.Unit, error) {
	var unit domain.Unit
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(code))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			unit = domain.Unit{
				CODE: code,
				URL:  string(val),
			}
			return nil
		})
		return err
	})
	return &unit, err
}
