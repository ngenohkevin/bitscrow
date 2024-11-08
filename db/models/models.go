// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Dispute struct {
	ID         uuid.UUID          `json:"id"`
	EscrowID   pgtype.UUID        `json:"escrow_id"`
	Reason     string             `json:"reason"`
	Resolved   pgtype.Bool        `json:"resolved"`
	Resolution pgtype.Text        `json:"resolution"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type EscrowTransaction struct {
	ID             uuid.UUID      `json:"id"`
	BuyerID        uuid.UUID      `json:"buyer_id"`
	SellerID       uuid.UUID      `json:"seller_id"`
	BitcoinAddress string         `json:"bitcoin_address"`
	Amount         pgtype.Numeric `json:"amount"`
	Status         pgtype.Text    `json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
}
