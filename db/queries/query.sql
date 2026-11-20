-- name: CreateUser :one
INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id; -- user_id が欲しいため

-- name: CreateToken :exec
INSERT INTO user_activations (user_id, token_hash, expires_at) VALUES ($1, $2, $3);