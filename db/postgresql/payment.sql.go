// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: payment.sql

package db

import (
	"context"
)

const createPayment = `-- name: CreatePayment :one

INSERT INTO payments (order_id, amount, payment_type, status)
VALUES ($1, $2, $3, $4)
RETURNING pid, order_id, amount, payment_type, status, created_at, updated_at
`

type CreatePaymentParams struct {
	OrderID     int64  `db:"order_id"`
	Amount      string `db:"amount"`
	PaymentType string `db:"payment_type"`
	Status      string `db:"status"`
}

// payments table
func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.OrderID,
		arg.Amount,
		arg.PaymentType,
		arg.Status,
	)
	var i Payment
	err := row.Scan(
		&i.Pid,
		&i.OrderID,
		&i.Amount,
		&i.PaymentType,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments WHERE pid = $1
`

func (q *Queries) DeletePayment(ctx context.Context, pid int64) error {
	_, err := q.db.ExecContext(ctx, deletePayment, pid)
	return err
}

const getAllPayments = `-- name: GetAllPayments :many
SELECT pid, order_id, amount, payment_type, status, created_at, updated_at FROM payments
`

func (q *Queries) GetAllPayments(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getAllPayments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.Pid,
			&i.OrderID,
			&i.Amount,
			&i.PaymentType,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPayment = `-- name: GetPayment :one
SELECT pid, order_id, amount, payment_type, status, created_at, updated_at FROM payments WHERE pid = $1
`

func (q *Queries) GetPayment(ctx context.Context, pid int64) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, pid)
	var i Payment
	err := row.Scan(
		&i.Pid,
		&i.OrderID,
		&i.Amount,
		&i.PaymentType,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePayment = `-- name: UpdatePayment :one
UPDATE payments
SET order_id = $2,
    amount = $3,
    payment_type = $4,
    status = $5,
    updated_at = now()
WHERE pid = $1
RETURNING pid, order_id, amount, payment_type, status, created_at, updated_at
`

type UpdatePaymentParams struct {
	Pid         int64  `db:"pid"`
	OrderID     int64  `db:"order_id"`
	Amount      string `db:"amount"`
	PaymentType string `db:"payment_type"`
	Status      string `db:"status"`
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, updatePayment,
		arg.Pid,
		arg.OrderID,
		arg.Amount,
		arg.PaymentType,
		arg.Status,
	)
	var i Payment
	err := row.Scan(
		&i.Pid,
		&i.OrderID,
		&i.Amount,
		&i.PaymentType,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}