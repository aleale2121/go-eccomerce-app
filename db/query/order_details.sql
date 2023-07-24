-- name: CreateOrderDetail :one
INSERT INTO order_details (user_id, total, payment_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetOrderDetail :one
SELECT * FROM order_details WHERE oid = $1;

-- name: GetAllOrderDetails :many
SELECT * FROM order_details;

-- name: UpdateOrderDetail :one
UPDATE order_details
SET user_id = $2,
    total = $3,
    payment_id = $4,
    updated_at = now()
WHERE oid = $1
RETURNING *;


-- name: DeleteOrderDetail :exec
DELETE FROM order_details WHERE oid = $1;

-- name: GetUserOrders :many
SELECT 
    json_agg(json_build_object(
        'order_id', od.oid,
        'total', od.total,
        'items', (
            SELECT json_agg(json_build_object(
                        'item_id', oi.oiid,
                        'product_id', oi.product_id,
                        'product_name', pr.name,
                        'product_description', pr.description,
                        'product_category', pr.category,
                        'product_price', pr.price,
                        'item_created_at', oi.created_at,
                        'item_updated_at', oi.updated_at
                    )
            ) FROM order_items oi
            JOIN products pr ON oi.product_id = pr.proid
            WHERE od.oid = oi.order_id
        ),
        'payment', json_build_object(
            'payment_id', p.pid,
            'amount', p.amount,
            'payment_type', p.payment_type,
            'status', p.status,
            'payment_created_at', p.created_at,
            'payment_updated_at', p.updated_at
        ),
        'order_created_at', od.created_at,
        'order_updated_at', od.updated_at
    ))
FROM 
    order_details od
JOIN
    payments p ON od.payment_id = p.pid
WHERE
    od.user_id = $1;
