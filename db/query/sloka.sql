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

-- name: AddSloka :one
INSERT INTO slokas (
  chapter_id,
  sloka,
  transliteration,
  purport,
  explanation
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateSloka :one
UPDATE slokas
SET
  sloka           = COALESCE($2, sloka),
  transliteration = COALESCE($3, transliteration),
  purport         = COALESCE($4, purport),
  explanation     = COALESCE($5, explanation)
WHERE id = $1
RETURNING *;