// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: escrow_transaction.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEscrowTransactions = `-- name: CreateEscrowTransactions :one
INSERT INTO escrow_transactions (id, buyer_id, seller_id, bitcoin_address, amount, status)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, buyer_id, seller_id, bitcoin_address, amount, status, created_at
`

type CreateEscrowTransactionsParams struct {
	ID             uuid.UUID      `json:"id"`
	BuyerID        uuid.UUID      `json:"buyer_id"`
	SellerID       uuid.UUID      `json:"seller_id"`
	BitcoinAddress string         `json:"bitcoin_address"`
	Amount         pgtype.Numeric `json:"amount"`
	Status         pgtype.Text    `json:"status"`
}

func (q *Queries) CreateEscrowTransactions(ctx context.Context, arg CreateEscrowTransactionsParams) (EscrowTransaction, error) {
	row := q.db.QueryRow(ctx, createEscrowTransactions,
		arg.ID,
		arg.BuyerID,
		arg.SellerID,
		arg.BitcoinAddress,
		arg.Amount,
		arg.Status,
	)
	var i EscrowTransaction
	err := row.Scan(
		&i.ID,
		&i.BuyerID,
		&i.SellerID,
		&i.BitcoinAddress,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getEscrowTransactions = `-- name: GetEscrowTransactions :one
SELECT id, buyer_id, seller_id, bitcoin_address, amount, status, created_at FROM escrow_transactions WHERE id = $1
`

func (q *Queries) GetEscrowTransactions(ctx context.Context, id uuid.UUID) (EscrowTransaction, error) {
	row := q.db.QueryRow(ctx, getEscrowTransactions, id)
	var i EscrowTransaction
	err := row.Scan(
		&i.ID,
		&i.BuyerID,
		&i.SellerID,
		&i.BitcoinAddress,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateEscrowTransactions = `-- name: UpdateEscrowTransactions :exec
UPDATE escrow_transactions
SET buyer_id = $2, seller_id = $3, bitcoin_address = $4, amount = $5, status = $6
WHERE id = $1
RETURNING id, buyer_id, seller_id, bitcoin_address, amount, status, created_at
`

type UpdateEscrowTransactionsParams struct {
	ID             uuid.UUID      `json:"id"`
	BuyerID        uuid.UUID      `json:"buyer_id"`
	SellerID       uuid.UUID      `json:"seller_id"`
	BitcoinAddress string         `json:"bitcoin_address"`
	Amount         pgtype.Numeric `json:"amount"`
	Status         pgtype.Text    `json:"status"`
}

func (q *Queries) UpdateEscrowTransactions(ctx context.Context, arg UpdateEscrowTransactionsParams) error {
	_, err := q.db.Exec(ctx, updateEscrowTransactions,
		arg.ID,
		arg.BuyerID,
		arg.SellerID,
		arg.BitcoinAddress,
		arg.Amount,
		arg.Status,
	)
	return err
}
