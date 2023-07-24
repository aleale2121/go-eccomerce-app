-- name: CreateProduct :one
INSERT INTO products (
    name, description, category, price, stock
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE proid = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY proid
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET name = $2,
    description = $3,
    category = $4,
    price = $5,
    stock = $6,
    updated_at = now()
WHERE proid = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE proid = $1;
