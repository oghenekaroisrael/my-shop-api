-- name: StatsForShops :many
SELECT * FROM banks
WHERE shop_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;


-- name: StatsForOneShop :many


