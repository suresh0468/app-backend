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
  title,
  subtitle,
  description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateChapter :one
UPDATE chapters
SET
  title       = COALESCE($2, title),
  subtitle    = COALESCE($3, subtitle),
  description = COALESCE($4, description)
WHERE id = $1
RETURNING *;