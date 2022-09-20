-- name: CreateShop :one
INSERT INTO shops (
  shop_name,
  shop_type,
  address,
  user_id
) VALUES (
  $1, $2, $3,$4
)
RETURNING *;

-- name: GetShopById :one
SELECT * FROM shops
WHERE id = $1 LIMIT 1;


-- name: ListMyShops :many
SELECT * FROM shops
WHERE user_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: UpdateShopDetail :one
UPDATE shops
SET 
    shop_name = $2,
    shop_type = $3,
    address = $4
WHERE id = $1
RETURNING *;

-- name: DeleteShop :exec
DELETE FROM shops
WHERE id = $1;

-- name: CountShops :one
SELECT COUNT(id) FROM shops
WHERE user_id = $1; 