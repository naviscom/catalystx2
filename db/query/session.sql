-- name: CreateSession :one
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at,
    created_at) VALUES (
 $1,$2,$3,$4,$5,$6,$7,$8
)
RETURNING *;

-- name: GetSession0 :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: ListSessions :many
SELECT * FROM sessions
WHERE username = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSession :one
UPDATE sessions
SET username = $2,
refresh_token = $3,
user_agent = $4,
client_ip = $5,
is_blocked = $6,
expires_at = $7,
created_at = $8
WHERE id = $1
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;
