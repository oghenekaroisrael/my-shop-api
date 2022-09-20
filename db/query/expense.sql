-- name: CreateExpenses :one
INSERT INTO expenses (
  particular,
  amount,
  recipient,
  payment_type,
  shop_id
) VALUES (
  $1, $2, $3,$4, $5
)
RETURNING *;

-- name: GetExpenseById :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;


-- name: ListMyExpenses :many
SELECT * FROM expenses
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: SearchExpensesBetweenDates :many
SELECT * FROM expenses
WHERE shop_id = $1 AND created_at BETWEEN $2 AND $3
ORDER BY id 
LIMIT $4
OFFSET $5;



-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;