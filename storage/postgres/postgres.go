package postgres

import (
	"context"
	"fmt"
	"uzum_orderclone/config"
	"uzum_orderclone/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db          *pgxpool.Pool
	order       storage.OrderyRepoI
	producOrder storage.OrderProductepoI
	orderstatus storage.OrderStatusI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Order() storage.OrderyRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) OrderProduct() storage.OrderProductepoI {
	if s.producOrder == nil {
		s.producOrder = NewOrderProductRepo(s.db)

	}
	return s.producOrder
}

func (s *Store) OrderStatus() storage.OrderStatusI {
	if s.orderstatus == nil {
		s.orderstatus = NewOrderStatuRepo(s.db)

	}
	return s.orderstatus
}
