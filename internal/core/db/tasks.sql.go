// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tasks.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (user_id, task_id)
VALUES ($1, $2)
RETURNING id, user_id, task_id, created_at, updated_at
`

type CreateTaskParams struct {
	UserID uuid.UUID
	TaskID string
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.UserID, arg.TaskID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
ORDER BY created_at DESC
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TaskID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasksByUserID = `-- name: ListTasksByUserID :many
SELECT id, user_id, task_id, created_at, updated_at
FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListTasksByUserID(ctx context.Context, userID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TaskID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks
SET task_id = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, user_id, task_id, created_at, updated_at
`

type UpdateTaskParams struct {
	ID     uuid.UUID
	TaskID string
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask, arg.ID, arg.TaskID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
