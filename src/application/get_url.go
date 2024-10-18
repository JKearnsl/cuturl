package application

import (
	"cuturl/src/application/common"
)

type GetUrlRequest struct {
	Code string
}

type GetUrlResponse struct {
	Url string
}

type GetUrl struct {
	UnitGateway common.UnitGateway
}

func (m *GetUrl) Execute(input *GetUrlRequest) (*GetUrlResponse, error) {

	url, err := m.UnitGateway.Get(input.Code)
	if err != nil {
		return nil, err
	}

	return &GetUrlResponse{
		Url: url.Url,
	}, nil
}
