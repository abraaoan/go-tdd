package provider

import (
	"fmt"
	"time"
)

type TokenProvider interface {
	Generate(userId string) (string, error)
}

type SimpleTokenProvider struct{}

func NewSimpleTokenProvider() *SimpleTokenProvider {
	return &SimpleTokenProvider{}
}

func (s *SimpleTokenProvider) Generate(userId string) (string, error) {
	token := fmt.Sprintf("token-%s-%d", userId, time.Now().UnixNano())
	return token, nil
}
