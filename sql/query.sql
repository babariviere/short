-- name: GetURLByHash :one
SELECT * FROM urls
WHERE hash = $1;
