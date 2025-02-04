-- name: CreateUser :one
INSERT INTO users (
    username,
    hash_password,
    full_name,
    email
) VALUES (
  $1, $2, $3, $4
)RETURNING *; -- we are returning the account id to the client

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 
LIMIT 1;
