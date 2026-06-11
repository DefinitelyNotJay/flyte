-- name: CreateLike :one
INSERT INTO likes (
    user_id,
    post_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetLike :one
SELECT * FROM likes
WHERE user_id = $1 AND post_id = $2 LIMIT 1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE user_id = $1 AND post_id = $2;
