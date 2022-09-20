-- name: CreatePayout :one
INSERT INTO payouts (
  particular,
  amount,
  recipient,
  payment_type,
  shop_id
) VALUES (
  $1, $2, $3,$4, $5
)
RETURNING *;

-- name: GetPayoutById :one
SELECT * FROM payouts
WHERE id = $1 LIMIT 1;


-- name: ListMyPayouts :many
SELECT * FROM payouts
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;


-- name: SearchBetweenDates :many
SELECT * FROM payouts
WHERE shop_id = $1 AND created_at BETWEEN $2 AND $3
ORDER BY id 
LIMIT $4
OFFSET $5;



-- name: DeletePayout :exec
DELETE FROM payouts
WHERE id = $1;