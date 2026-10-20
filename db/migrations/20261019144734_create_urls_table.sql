-- +goose Up
CREATE TABLE urls (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL,
    short_code TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE urls;
