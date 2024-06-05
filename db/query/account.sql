-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccount :many
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: GetAccountList :many
SELECT *
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :exec
UPDATE accounts
SET balance = $1
WHERE id = $2;

-- name: UpdateAccountWithReturn :one
UPDATE accounts
SET balance = $1
WHERE id = $2
RETURNING *;