package service

import (
	"encoding/json"
	"errors"
	"exchange-rate-receiver/internal/models"
	"exchange-rate-receiver/internal/repository"
	"net/http"
	"time"
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
	rates, err := fetchRatesFromGarantex()
	if err != nil {
		return nil, err
	}

	rates.Timestamp = time.Now().Format(time.RFC3339)
	err = s.repo.StoreRates(rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func fetchRatesFromGarantex() (*models.Rates, error) {
	resp, err := http.Get("https://garantex.org/api/v2/depth")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch rates from Garantex")
	}

	var data struct {
		Asks [][]interface{} `json:"asks"`
		Bids [][]interface{} `json:"bids"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Asks) == 0 || len(data.Bids) == 0 {
		return nil, errors.New("no ask or bid prices available")
	}

	ask, ok1 := data.Asks[0][0].(float64)
	bid, ok2 := data.Bids[0][0].(float64)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid data format from Garantex")
	}

	return &models.Rates{
		Ask: ask,
		Bid: bid,
	}, nil
}
