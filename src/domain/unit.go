package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Unit struct {
	Code string `json:"code"`
	Url  string `json:"url"`
}

func CreateUnit(url string) (*Unit, error) {
	if url == "" {
		return nil, errors.New("code and url are required")
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return &Unit{
		Code: string(code),
		Url:  url,
	}, nil
}
