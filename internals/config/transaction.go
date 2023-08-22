package config

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

var ctxKey = struct{}{}

type tx struct {
	pc *PostgresClient
}

var ErrNoTransaction = errors.New("no transaction in context")

// Transaction トランザクション
type Transaction interface {
	RunInTx(context.Context, func(context.Context) error) error
}

// NewTx トランザクション管理
func NewTx(db *PostgresClient) Transaction {
	return &tx{db}
}

func (t *tx) RunInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.pc.Write.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		ctx = context.WithValue(ctx, ctxKey, tx)
		return fn(ctx)
	})
}

// GetTx トランザクション取得
func GetTx(ctx context.Context) (bun.IDB, error) {
	tx, ok := ctx.Value(ctxKey).(bun.IDB)
	if !ok {
		return nil, ErrNoTransaction
	}
	return tx, nil
}
