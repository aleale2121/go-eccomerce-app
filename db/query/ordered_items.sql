-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetOrderItemByID :one
SELECT * FROM order_items WHERE oiid = $1;

-- name: GetOrderItemByOrderID :many
SELECT * FROM order_items
WHERE order_id = $1;

-- name: GetOrderItemByProductID :many
SELECT * FROM order_items
WHERE product_id = $1;;

-- name: UpdateOrderItem :one
UPDATE order_items
SET order_id = $2,
    product_id = $3,
    updated_at = now()
WHERE oiid = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE oiid = $1;