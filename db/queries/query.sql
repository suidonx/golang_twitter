-- name: CreateUser :one
INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id; -- user_id が欲しいため

-- name: CreateToken :exec
INSERT INTO user_activations (user_id, token_hash, expires_at) VALUES ($1, $2, $3);

-- name: FindEmailActivationByTokenHash :one
SELECT user_id, token_hash, expires_at
FROM user_activations
WHERE token_hash = $1;

-- name: MarkEmailAsVerified :exec
UPDATE users
SET email_verified_at = NOW()
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, password_hash, email_verified_at
FROM users
WHERE email = $1;

-- name: CreatePost :exec
INSERT INTO posts (content, user_id) VALUES ($1, $2);

