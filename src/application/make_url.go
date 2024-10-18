package application

import (
	"cuturl/src/application/common"
	"cuturl/src/domain"
)

type MakeUrlRequest struct {
	Url string
}

type MakeUrlResponse struct {
	Code string
}

type MakeUrl struct {
	UnitGateway common.UnitGateway
}

func (m *MakeUrl) Execute(input *MakeUrlRequest) (*MakeUrlResponse, error) {

	unit, err := domain.CreateUnit(input.Url)
	if err != nil {
		return nil, err
	}

	err = m.UnitGateway.Save(unit)
	if err != nil {
		return nil, err
	}

	return &MakeUrlResponse{
		Code: unit.Code,
	}, nil
}
