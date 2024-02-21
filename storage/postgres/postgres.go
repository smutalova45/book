package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/config"
	"main.go/storage"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	// //migration
	m, err := migrate.New("file://migrations/postgres/", url)
	if err != nil {
		return nil, err
	}

	if err = m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			if err != nil {
				return nil, err
			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					return nil, err
				}
			}
			return nil, err
		}
	}

	return Store{
		pool: pool,
	}, nil
}
func (s Store) Close() {
	s.pool.Close()

}
func (s Store) Book() storage.IBookStorage {
	return NewBookRepo(s.pool)
}
