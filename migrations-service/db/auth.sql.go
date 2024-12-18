// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auth.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAuthProvider = `-- name: CreateAuthProvider :one
INSERT INTO user_auth_providers (user_id, provider, provider_user_id, access_token, refresh_token, token_expiry)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, user_id, provider, provider_user_id, created_at, updated_at
`

type CreateAuthProviderParams struct {
	UserID         uuid.UUID
	Provider       string
	ProviderUserID string
	AccessToken    sql.NullString
	RefreshToken   sql.NullString
	TokenExpiry    sql.NullTime
}

type CreateAuthProviderRow struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Provider       string
	ProviderUserID string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) CreateAuthProvider(ctx context.Context, arg CreateAuthProviderParams) (CreateAuthProviderRow, error) {
	row := q.db.QueryRowContext(ctx, createAuthProvider,
		arg.UserID,
		arg.Provider,
		arg.ProviderUserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.TokenExpiry,
	)
	var i CreateAuthProviderRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.ProviderUserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAuthProviderByProviderUserId = `-- name: GetAuthProviderByProviderUserId :one
SELECT id, user_id, provider, provider_user_id, access_token, refresh_token, token_expiry, created_at, updated_at
FROM user_auth_providers
WHERE provider_user_id = $1
`

func (q *Queries) GetAuthProviderByProviderUserId(ctx context.Context, providerUserID string) (UserAuthProvider, error) {
	row := q.db.QueryRowContext(ctx, getAuthProviderByProviderUserId, providerUserID)
	var i UserAuthProvider
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.ProviderUserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.TokenExpiry,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
