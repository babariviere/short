-- name: GetURLByHash :one
SELECT * FROM urls
WHERE hash = $1;

-- name: InsertURL :one
INSERT INTO urls (hash, long_url)
VALUES ($1, $2)
ON CONFLICT (hash) DO UPDATE SET hash = EXCLUDED.hash -- force row to be returned
RETURNING *;
