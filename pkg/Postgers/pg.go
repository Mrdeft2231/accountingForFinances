package Postgers

import (
	"accounting/config"
	"context"
	"errors"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func InitDB(ctx context.Context, cfg config.Postgres, logger *zap.Logger) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.URL)
	if err != nil {
		logger.Panic("Не удалось подключиться к БД", zap.Error(err))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		logger.Panic("Не удалось отправить ping", zap.Error(err))
		return nil, err
	}

	m, err := migrate.New("file://../../migrations", cfg.URL)
	if err != nil {
		pool.Close()
		logger.Error("не удалось создать миграции", zap.Error(err))
		return nil, err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		pool.Close()
		logger.Error("Не удалось мигрировать данные", zap.Error(err))
		return nil, err
	}

	return pool, nil
}
