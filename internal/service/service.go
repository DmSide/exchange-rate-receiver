package service

import (
	"exchange-rate-receiver/internal/models"
	"exchange-rate-receiver/internal/repository"
)

type Service interface {
	GetRates() (*models.Rates, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetRates() (*models.Rates, error) {
	// TODO: Implement
	return &models.Rates{Ask: 1, Bid: 1}, nil
}
