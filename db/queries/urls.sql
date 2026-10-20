-- name: CreateUrl :one
INSERT INTO urls (
    id,
    created_at,
    updated_at,
    url,
    short_code
) values (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetUrlByShortCode :one
SELECT * FROM urls WHERE short_code = $1;
