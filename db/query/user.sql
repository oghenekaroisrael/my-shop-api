-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT * FROM users
WHERE email = $1 AND password = $2 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id 
LIMIT $1 
OFFSET $2;

-- name: UpdateUserVerificationStatus :one
UPDATE users
SET isVerified = true
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
set password = $3
WHERE 
        id = $1  OR
        email = $2  
RETURNING *;

-- name: UpdateUserDetail :one
UPDATE users
set 
    first_name = $2,
    last_name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;