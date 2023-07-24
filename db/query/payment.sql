-- payments table

-- name: CreatePayment :one
INSERT INTO payments (order_id, amount, payment_type, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE pid = $1;

-- name: GetAllPayments :many
SELECT * FROM payments;

-- name: UpdatePayment :one
UPDATE payments
SET order_id = $2,
    amount = $3,
    payment_type = $4,
    status = $5,
    updated_at = now()
WHERE pid = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments WHERE pid = $1;

