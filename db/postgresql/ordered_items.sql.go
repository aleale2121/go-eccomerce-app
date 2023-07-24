// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: ordered_items.sql

package db

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id)
VALUES ($1, $2)
RETURNING oiid, order_id, product_id, created_at, updated_at
`

type CreateOrderItemParams struct {
	OrderID   int64 `db:"order_id"`
	ProductID int64 `db:"product_id"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem, arg.OrderID, arg.ProductID)
	var i OrderItem
	err := row.Scan(
		&i.Oiid,
		&i.OrderID,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE oiid = $1
`

func (q *Queries) DeleteOrderItem(ctx context.Context, oiid int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItem, oiid)
	return err
}

const getOrderItemByID = `-- name: GetOrderItemByID :one
SELECT oiid, order_id, product_id, created_at, updated_at FROM order_items WHERE oiid = $1
`

func (q *Queries) GetOrderItemByID(ctx context.Context, oiid int64) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderItemByID, oiid)
	var i OrderItem
	err := row.Scan(
		&i.Oiid,
		&i.OrderID,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderItemByOrderID = `-- name: GetOrderItemByOrderID :many
SELECT oiid, order_id, product_id, created_at, updated_at FROM order_items
WHERE order_id = $1
`

func (q *Queries) GetOrderItemByOrderID(ctx context.Context, orderID int64) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getOrderItemByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.Oiid,
			&i.OrderID,
			&i.ProductID,
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

const getOrderItemByProductID = `-- name: GetOrderItemByProductID :many
SELECT oiid, order_id, product_id, created_at, updated_at FROM order_items
WHERE product_id = $1
`

func (q *Queries) GetOrderItemByProductID(ctx context.Context, productID int64) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getOrderItemByProductID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.Oiid,
			&i.OrderID,
			&i.ProductID,
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

const updateOrderItem = `-- name: UpdateOrderItem :one
UPDATE order_items
SET order_id = $2,
    product_id = $3,
    updated_at = now()
WHERE oiid = $1
RETURNING oiid, order_id, product_id, created_at, updated_at
`

type UpdateOrderItemParams struct {
	Oiid      int64 `db:"oiid"`
	OrderID   int64 `db:"order_id"`
	ProductID int64 `db:"product_id"`
}

func (q *Queries) UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, updateOrderItem, arg.Oiid, arg.OrderID, arg.ProductID)
	var i OrderItem
	err := row.Scan(
		&i.Oiid,
		&i.OrderID,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
