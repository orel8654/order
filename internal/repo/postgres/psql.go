package postgres

import (
	"context"
	"order/internal/types"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateOrder(ctx context.Context, data types.CreateOrderUUID) error {
	return nil
}

func (r *Repo) GetMenu(ctx context.Context) (types.Menu, error) {
	return types.Menu{}, nil
}
