ALTER TABLE chapters
DROP COLUMN description;

ALTER TABLE chapters
DROP COLUMN subtitle;

ALTER TABLE chapters
RENAME COLUMN title TO name;