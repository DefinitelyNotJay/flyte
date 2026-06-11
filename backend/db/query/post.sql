-- name: CreatePost :one
INSERT INTO posts (
    author_id,
    content,
    parent_id,
    repost_of_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET
    content = COALESCE(sqlc.narg(content), content),
    reply_count = COALESCE(sqlc.narg(reply_count), reply_count),
    like_count = COALESCE(sqlc.narg(like_count), like_count),
    repost_count = COALESCE(sqlc.narg(repost_count), repost_count),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
UPDATE posts
SET deleted_at = now()
WHERE id = $1;
