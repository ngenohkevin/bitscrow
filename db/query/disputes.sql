-- name: CreateDisputes :one
INSERT INTO disputes(id, escrow_id, reason, resolved, resolution, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetDisputes :one
SELECT * FROM disputes WHERE id = $1;

-- name: UpdateDisputes :exec
UPDATE disputes
SET escrow_id = $2, reason = $3, resolved = $4, resolution = $5, created_at = $6, updated_at = $7
WHERE id = $1
RETURNING *;