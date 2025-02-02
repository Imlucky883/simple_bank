-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
  $1, $2, $3
)RETURNING *; -- we are returning the account id to the client

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- Example Usage:
-- If you execute this query with parameters 10{$1} and 20{$2}, it will return 10 rows starting 
-- from the 21st row (since the offset is 20).