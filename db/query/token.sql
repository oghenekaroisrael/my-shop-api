-- name: CreateToken :one
INSERT INTO tokens (
  user_id,
  access_token,
  refresh_token
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserRefresToken :one
SELECT user_id, refresh_token FROM tokens
WHERE user_id = $1
LIMIT 1;

-- name: GetUserAccessToken :one
SELECT user_id, access_token FROM tokens
WHERE user_id = $1
LIMIT 1;
