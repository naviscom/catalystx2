// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: country.sql

package db

import (
	"context"
)

const createCountry = `-- name: CreateCountry :one
INSERT INTO countries (
    country_name,
    country_desc,
    continent_id) VALUES (
 $1,$2,$3
)
RETURNING id, country_name, country_desc, continent_id
`

type CreateCountryParams struct {
	CountryName string `json:"country_name"`
	CountryDesc string `json:"country_desc"`
	ContinentID int64  `json:"continent_id"`
}

func (q *Queries) CreateCountry(ctx context.Context, arg CreateCountryParams) (Country, error) {
	row := q.db.QueryRow(ctx, createCountry, arg.CountryName, arg.CountryDesc, arg.ContinentID)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryDesc,
		&i.ContinentID,
	)
	return i, err
}

const deleteCountry = `-- name: DeleteCountry :exec
DELETE FROM countries
WHERE id = $1
`

func (q *Queries) DeleteCountry(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCountry, id)
	return err
}

const getCountry0 = `-- name: GetCountry0 :one
SELECT id, country_name, country_desc, continent_id FROM countries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCountry0(ctx context.Context, id int64) (Country, error) {
	row := q.db.QueryRow(ctx, getCountry0, id)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryDesc,
		&i.ContinentID,
	)
	return i, err
}

const getCountry1 = `-- name: GetCountry1 :one
SELECT id, country_name, country_desc, continent_id FROM countries
WHERE country_name = $1 LIMIT 1
`

func (q *Queries) GetCountry1(ctx context.Context, countryName string) (Country, error) {
	row := q.db.QueryRow(ctx, getCountry1, countryName)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryDesc,
		&i.ContinentID,
	)
	return i, err
}

const listCountries = `-- name: ListCountries :many
SELECT id, country_name, country_desc, continent_id FROM countries
WHERE continent_id = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCountriesParams struct {
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
	ContinentID int64 `json:"continent_id"`
}

func (q *Queries) ListCountries(ctx context.Context, arg ListCountriesParams) ([]Country, error) {
	rows, err := q.db.Query(ctx, listCountries, arg.Limit, arg.Offset, arg.ContinentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Country{}
	for rows.Next() {
		var i Country
		if err := rows.Scan(
			&i.ID,
			&i.CountryName,
			&i.CountryDesc,
			&i.ContinentID,
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

const updateCountry = `-- name: UpdateCountry :one
UPDATE countries
SET country_name = $2,
country_desc = $3,
continent_id = $4
WHERE id = $1
RETURNING id, country_name, country_desc, continent_id
`

type UpdateCountryParams struct {
	ID          int64  `json:"id"`
	CountryName string `json:"country_name"`
	CountryDesc string `json:"country_desc"`
	ContinentID int64  `json:"continent_id"`
}

func (q *Queries) UpdateCountry(ctx context.Context, arg UpdateCountryParams) (Country, error) {
	row := q.db.QueryRow(ctx, updateCountry,
		arg.ID,
		arg.CountryName,
		arg.CountryDesc,
		arg.ContinentID,
	)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryDesc,
		&i.ContinentID,
	)
	return i, err
}
