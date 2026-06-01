-- name: GetSloka :one
SELECT *
FROM slokas
WHERE id = $1
LIMIT 1;

-- name: ListSlokasByChapter :many
SELECT *
FROM slokas
WHERE chapter_id = $1
ORDER BY id;