CREATE TABLE chapters (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);


CREATE TABLE slokas (
    id BIGSERIAL PRIMARY KEY,
    chapter_id BIGINT NOT NULL,
    sloka TEXT NOT NULL,
    transliteration TEXT NOT NULL,
    purport TEXT,
    explanation TEXT,

    CONSTRAINT fk_chapter
        FOREIGN KEY (chapter_id)
        REFERENCES chapters(id)
);