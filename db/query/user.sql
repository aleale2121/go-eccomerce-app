-- name: CreateUser :one
INSERT INTO users (
  username, password, name, address, mobile_no
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE uid = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY uid
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
  set username = $2,
  password = $3,
  name = $4,
  address = $5,
  mobile_no = $6
WHERE uid = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE uid = $1;
