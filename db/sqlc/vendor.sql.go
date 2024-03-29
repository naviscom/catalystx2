// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: vendor.sql

package db

import (
	"context"
)

const createVendor = `-- name: CreateVendor :one
INSERT INTO vendors (
    vendor_name,
    vendor_desc) VALUES (
 $1,$2
)
RETURNING id, vendor_name, vendor_desc
`

type CreateVendorParams struct {
	VendorName string `json:"vendor_name"`
	VendorDesc string `json:"vendor_desc"`
}

func (q *Queries) CreateVendor(ctx context.Context, arg CreateVendorParams) (Vendor, error) {
	row := q.db.QueryRow(ctx, createVendor, arg.VendorName, arg.VendorDesc)
	var i Vendor
	err := row.Scan(&i.ID, &i.VendorName, &i.VendorDesc)
	return i, err
}

const deleteVendor = `-- name: DeleteVendor :exec
DELETE FROM vendors
WHERE id = $1
`

func (q *Queries) DeleteVendor(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteVendor, id)
	return err
}

const getVendor0 = `-- name: GetVendor0 :one
SELECT id, vendor_name, vendor_desc FROM vendors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVendor0(ctx context.Context, id int64) (Vendor, error) {
	row := q.db.QueryRow(ctx, getVendor0, id)
	var i Vendor
	err := row.Scan(&i.ID, &i.VendorName, &i.VendorDesc)
	return i, err
}

const getVendor1 = `-- name: GetVendor1 :one
SELECT id, vendor_name, vendor_desc FROM vendors
WHERE vendor_name = $1 LIMIT 1
`

func (q *Queries) GetVendor1(ctx context.Context, vendorName string) (Vendor, error) {
	row := q.db.QueryRow(ctx, getVendor1, vendorName)
	var i Vendor
	err := row.Scan(&i.ID, &i.VendorName, &i.VendorDesc)
	return i, err
}

const listVendors = `-- name: ListVendors :many
SELECT id, vendor_name, vendor_desc FROM vendors
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListVendorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListVendors(ctx context.Context, arg ListVendorsParams) ([]Vendor, error) {
	rows, err := q.db.Query(ctx, listVendors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vendor{}
	for rows.Next() {
		var i Vendor
		if err := rows.Scan(&i.ID, &i.VendorName, &i.VendorDesc); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateVendor = `-- name: UpdateVendor :one
UPDATE vendors
SET vendor_name = $2,
vendor_desc = $3
WHERE id = $1
RETURNING id, vendor_name, vendor_desc
`

type UpdateVendorParams struct {
	ID         int64  `json:"id"`
	VendorName string `json:"vendor_name"`
	VendorDesc string `json:"vendor_desc"`
}

func (q *Queries) UpdateVendor(ctx context.Context, arg UpdateVendorParams) (Vendor, error) {
	row := q.db.QueryRow(ctx, updateVendor, arg.ID, arg.VendorName, arg.VendorDesc)
	var i Vendor
	err := row.Scan(&i.ID, &i.VendorName, &i.VendorDesc)
	return i, err
}
