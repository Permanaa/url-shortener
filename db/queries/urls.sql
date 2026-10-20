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

-- name: UpdateUrlByShortCode :one
UPDATE urls
SET url = $1, updated_at = NOW()
WHERE short_code = $2
RETURNING *;

-- name: DeleteUrlByShortCode :exec
DELETE FROM urls WHERE short_code = $1;

-- name: IncrementAccessCount :exec
UPDATE urls
SET access_count = access_count + 1
WHERE short_code = $1;
