-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY name;

-- name: CreateEvent :one
INSERT INTO events (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateEvent :exec
UPDATE events
  set name = $2
WHERE id = $1;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;