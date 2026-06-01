-- name: GetChapter :one
SELECT *
FROM chapters
WHERE id = $1
LIMIT 1;

-- name: ListChapters :many
SELECT *
FROM chapters
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddChapter :one
INSERT INTO chapters (
  name
) VALUES (
  $1
)
RETURNING *;