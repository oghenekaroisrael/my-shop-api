-- name: CreateItem :one
INSERT INTO inventories (
  item_name,
  quantity,
  cost_price,
  selling_price_standard,
  status,
  shop_id
) VALUES (
  $1, $2, $3,$4, $5, $6
)
RETURNING *;

-- name: GetItemById :one
SELECT * FROM inventories
WHERE id = $1 LIMIT 1;


-- name: ListMyInventories :many
SELECT * FROM inventories
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: SearchInventoryByName :many
SELECT * FROM inventories
WHERE shop_id = $1 AND item_name = $2
ORDER BY id 
LIMIT $3
OFFSET $4;

-- name: UpdateItemDetail :one
UPDATE inventories
set 
    item_name = $2,
    cost_price = $3,
    selling_price_standard = $4,
    status = $5
WHERE id = $1
RETURNING *;

-- name: UpdateItemQuantity :one
UPDATE inventories
set 
    quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM inventories
WHERE id = $1;