package services

import (
	"context"

	"github.com/google/uuid"
)

type UUIDIdentityService struct{}

func NewUUIDIdentityService() *UUIDIdentityService {
	return &UUIDIdentityService{}
}

func (s *UUIDIdentityService) Generate(ctx context.Context) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
