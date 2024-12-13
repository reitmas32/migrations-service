-- name: CreateAuthProvider :one
INSERT INTO user_auth_providers (user_id, provider, provider_user_id, access_token, refresh_token, token_expiry)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, user_id, provider, provider_user_id, created_at, updated_at;

-- name: GetAuthProviderByProviderUserId :one
SELECT id, user_id, provider, provider_user_id, access_token, refresh_token, token_expiry, created_at, updated_at
FROM user_auth_providers
WHERE provider_user_id = $1;