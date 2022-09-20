-- name: CreateBank :one
INSERT INTO banks (
  bank_name,
  icon,
  account_number,
  shop_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateBank :one
UPDATE banks
SET bank_name = $1,
    account_number = $2,
    icon = $3
WHERE id = $4
RETURNING *;

-- name: GetBankById :one
SELECT * FROM banks
WHERE id = $1 LIMIT 1;


-- name: ListMyBanks :many
SELECT * FROM banks
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: DeleteBank :exec
DELETE FROM banks
WHERE id = $1;