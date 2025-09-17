-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name AS feed_name, f.url AS feed_url, u.name AS username
FROM feeds f INNER JOIN users u
    ON f.user_id = u.id
;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1;