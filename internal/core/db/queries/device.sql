-- name: CreateDevice :one
INSERT INTO devices (operative_system, os_version, token, model)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDeviceByToken :one
SELECT *
FROM devices
WHERE token = $1;
