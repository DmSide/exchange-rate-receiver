package repository

import (
	"context"
	"exchange-rate-receiver/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	StoreRates(rates *models.Rates) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) StoreRates(rates *models.Rates) error {
	query := "INSERT INTO rates (ask, bid, timestamp) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(context.Background(), query, rates.Ask, rates.Bid, rates.Timestamp)
	return err
}
