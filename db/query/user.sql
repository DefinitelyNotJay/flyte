-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash,
    display_name,
    bio,
    avatar_url
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
    username = $2,
    email = $3,
    password_hash = $4,
    display_name = $5,
    bio = $6,
    avatar_url = $7,
    updated_at = now()
WHERE id = $1
RETURNING *;