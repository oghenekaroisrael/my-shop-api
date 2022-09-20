-- name: CreateSale :one
INSERT INTO sales (
  item_id,
  quantity,
  selling_price_actual,
  payment_type,
  shop_id
) VALUES (
  $1, $2, $3,$4, $5
)
RETURNING *;

-- name: GetSaleById :one
SELECT * FROM sales
WHERE id = $1 LIMIT 1;


-- name: ListMySales :many
SELECT * FROM sales
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: ListSalesByUser :many
SELECT 
sale.id AS id, 
inventories.item_name AS item_name,
sale.item_id AS item_id, 
sale.quantity AS quantity, 
sale.selling_price_actual AS selling_price,
sale.payment_type AS payment_type,
sale.created_at AS created_at
FROM sales sale
INNER JOIN shops ON sale.shop_id = shops.id
INNER JOIN users ON shops.user_id = users.id
INNER JOIN inventories ON sale.item_id = inventories.id
WHERE users.id = $1
ORDER BY sale.created_at DESC
LIMIT $2
OFFSET $3;

-- name: SearchSaleBetweenDates :many
SELECT * FROM sales
WHERE shop_id = $1 AND created_at BETWEEN $2 AND $2
ORDER BY id 
LIMIT $3
OFFSET $4;



-- name: DeleteSale :exec
DELETE FROM sales
WHERE id = $1;