package common

import (
	"cuturl/src/domain"
)

type UnitGateway interface {
	Save(unit *domain.Unit) error
	Get(code string) (*domain.Unit, error)
}
