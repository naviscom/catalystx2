// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: carrier.sql

package db

import (
	"context"
)

const createCarrier = `-- name: CreateCarrier :one
INSERT INTO carriers (
    carrier_name,
    carrier_desc,
    size,
    start_freq,
    end_freq,
    band_id) VALUES (
 $1,$2,$3,$4,$5,$6
)
RETURNING id, carrier_name, carrier_desc, size, start_freq, end_freq, band_id
`

type CreateCarrierParams struct {
	CarrierName string `json:"carrier_name"`
	CarrierDesc string `json:"carrier_desc"`
	Size        int64  `json:"size"`
	StartFreq   int64  `json:"start_freq"`
	EndFreq     int64  `json:"end_freq"`
	BandID      int64  `json:"band_id"`
}

func (q *Queries) CreateCarrier(ctx context.Context, arg CreateCarrierParams) (Carrier, error) {
	row := q.db.QueryRow(ctx, createCarrier,
		arg.CarrierName,
		arg.CarrierDesc,
		arg.Size,
		arg.StartFreq,
		arg.EndFreq,
		arg.BandID,
	)
	var i Carrier
	err := row.Scan(
		&i.ID,
		&i.CarrierName,
		&i.CarrierDesc,
		&i.Size,
		&i.StartFreq,
		&i.EndFreq,
		&i.BandID,
	)
	return i, err
}

const deleteCarrier = `-- name: DeleteCarrier :exec
DELETE FROM carriers
WHERE id = $1
`

func (q *Queries) DeleteCarrier(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCarrier, id)
	return err
}

const getCarrier0 = `-- name: GetCarrier0 :one
SELECT id, carrier_name, carrier_desc, size, start_freq, end_freq, band_id FROM carriers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCarrier0(ctx context.Context, id int64) (Carrier, error) {
	row := q.db.QueryRow(ctx, getCarrier0, id)
	var i Carrier
	err := row.Scan(
		&i.ID,
		&i.CarrierName,
		&i.CarrierDesc,
		&i.Size,
		&i.StartFreq,
		&i.EndFreq,
		&i.BandID,
	)
	return i, err
}

const getCarrier1 = `-- name: GetCarrier1 :one
SELECT id, carrier_name, carrier_desc, size, start_freq, end_freq, band_id FROM carriers
WHERE carrier_name = $1 LIMIT 1
`

func (q *Queries) GetCarrier1(ctx context.Context, carrierName string) (Carrier, error) {
	row := q.db.QueryRow(ctx, getCarrier1, carrierName)
	var i Carrier
	err := row.Scan(
		&i.ID,
		&i.CarrierName,
		&i.CarrierDesc,
		&i.Size,
		&i.StartFreq,
		&i.EndFreq,
		&i.BandID,
	)
	return i, err
}

const listCarriers = `-- name: ListCarriers :many
SELECT id, carrier_name, carrier_desc, size, start_freq, end_freq, band_id FROM carriers
WHERE band_id = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCarriersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	BandID int64 `json:"band_id"`
}

func (q *Queries) ListCarriers(ctx context.Context, arg ListCarriersParams) ([]Carrier, error) {
	rows, err := q.db.Query(ctx, listCarriers, arg.Limit, arg.Offset, arg.BandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Carrier{}
	for rows.Next() {
		var i Carrier
		if err := rows.Scan(
			&i.ID,
			&i.CarrierName,
			&i.CarrierDesc,
			&i.Size,
			&i.StartFreq,
			&i.EndFreq,
			&i.BandID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCarrier = `-- name: UpdateCarrier :one
UPDATE carriers
SET carrier_name = $2,
carrier_desc = $3,
size = $4,
start_freq = $5,
end_freq = $6,
band_id = $7
WHERE id = $1
RETURNING id, carrier_name, carrier_desc, size, start_freq, end_freq, band_id
`

type UpdateCarrierParams struct {
	ID          int64  `json:"id"`
	CarrierName string `json:"carrier_name"`
	CarrierDesc string `json:"carrier_desc"`
	Size        int64  `json:"size"`
	StartFreq   int64  `json:"start_freq"`
	EndFreq     int64  `json:"end_freq"`
	BandID      int64  `json:"band_id"`
}

func (q *Queries) UpdateCarrier(ctx context.Context, arg UpdateCarrierParams) (Carrier, error) {
	row := q.db.QueryRow(ctx, updateCarrier,
		arg.ID,
		arg.CarrierName,
		arg.CarrierDesc,
		arg.Size,
		arg.StartFreq,
		arg.EndFreq,
		arg.BandID,
	)
	var i Carrier
	err := row.Scan(
		&i.ID,
		&i.CarrierName,
		&i.CarrierDesc,
		&i.Size,
		&i.StartFreq,
		&i.EndFreq,
		&i.BandID,
	)
	return i, err
}
