-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING *;

-- name: GetEntry :many
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;

-- name: GetEntryList :many
SELECT *
FROM entries
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateEntry :exec
UPDATE entries
SET amount = $1
WHERE id = $2;

-- name: UpdateEntryWithReturn :one
UPDATE entries
SET amount = $1
WHERE id = $2
RETURNING *;

-- name: DeleteEntry :exec
DELETE
FROM entries
where id = $1;