-- name: CreateEscrowTransactions :one
INSERT INTO escrow_transactions (id, buyer_id, seller_id, bitcoin_address, amount, status)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetEscrowTransactions :one
SELECT * FROM escrow_transactions WHERE id = $1;

-- name: UpdateEscrowTransactions :exec
UPDATE escrow_transactions
SET buyer_id = $2, seller_id = $3, bitcoin_address = $4, amount = $5, status = $6
WHERE id = $1
RETURNING *;