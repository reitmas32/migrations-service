-- name: CreateTask :one
INSERT INTO tasks (user_id, task_id)
VALUES ($1, $2)
RETURNING id, user_id, task_id, created_at, updated_at;

-- name: GetTaskByID :one
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
WHERE id = $1;

-- name: ListTasks :many
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET task_id = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, user_id, task_id, created_at, updated_at;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: ListTasksByUserID :many
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC;
