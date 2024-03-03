// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: district.sql

package db

import (
	"context"
)

const createDistrict = `-- name: CreateDistrict :one
INSERT INTO districts (
    district_name,
    district_desc,
    city_id) VALUES (
 $1,$2,$3
)
RETURNING id, district_name, district_desc, city_id
`

type CreateDistrictParams struct {
	DistrictName string `json:"district_name"`
	DistrictDesc string `json:"district_desc"`
	CityID       int64  `json:"city_id"`
}

func (q *Queries) CreateDistrict(ctx context.Context, arg CreateDistrictParams) (District, error) {
	row := q.db.QueryRow(ctx, createDistrict, arg.DistrictName, arg.DistrictDesc, arg.CityID)
	var i District
	err := row.Scan(
		&i.ID,
		&i.DistrictName,
		&i.DistrictDesc,
		&i.CityID,
	)
	return i, err
}

const deleteDistrict = `-- name: DeleteDistrict :exec
DELETE FROM districts
WHERE id = $1
`

func (q *Queries) DeleteDistrict(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteDistrict, id)
	return err
}

const getDistrict0 = `-- name: GetDistrict0 :one
SELECT id, district_name, district_desc, city_id FROM districts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDistrict0(ctx context.Context, id int64) (District, error) {
	row := q.db.QueryRow(ctx, getDistrict0, id)
	var i District
	err := row.Scan(
		&i.ID,
		&i.DistrictName,
		&i.DistrictDesc,
		&i.CityID,
	)
	return i, err
}

const getDistrict1 = `-- name: GetDistrict1 :one
SELECT id, district_name, district_desc, city_id FROM districts
WHERE district_name = $1 LIMIT 1
`

func (q *Queries) GetDistrict1(ctx context.Context, districtName string) (District, error) {
	row := q.db.QueryRow(ctx, getDistrict1, districtName)
	var i District
	err := row.Scan(
		&i.ID,
		&i.DistrictName,
		&i.DistrictDesc,
		&i.CityID,
	)
	return i, err
}

const listDistricts = `-- name: ListDistricts :many
SELECT id, district_name, district_desc, city_id FROM districts
WHERE city_id = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListDistrictsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	CityID int64 `json:"city_id"`
}

func (q *Queries) ListDistricts(ctx context.Context, arg ListDistrictsParams) ([]District, error) {
	rows, err := q.db.Query(ctx, listDistricts, arg.Limit, arg.Offset, arg.CityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []District{}
	for rows.Next() {
		var i District
		if err := rows.Scan(
			&i.ID,
			&i.DistrictName,
			&i.DistrictDesc,
			&i.CityID,
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

const updateDistrict = `-- name: UpdateDistrict :one
UPDATE districts
SET district_name = $2,
district_desc = $3,
city_id = $4
WHERE id = $1
RETURNING id, district_name, district_desc, city_id
`

type UpdateDistrictParams struct {
	ID           int64  `json:"id"`
	DistrictName string `json:"district_name"`
	DistrictDesc string `json:"district_desc"`
	CityID       int64  `json:"city_id"`
}

func (q *Queries) UpdateDistrict(ctx context.Context, arg UpdateDistrictParams) (District, error) {
	row := q.db.QueryRow(ctx, updateDistrict,
		arg.ID,
		arg.DistrictName,
		arg.DistrictDesc,
		arg.CityID,
	)
	var i District
	err := row.Scan(
		&i.ID,
		&i.DistrictName,
		&i.DistrictDesc,
		&i.CityID,
	)
	return i, err
}
