// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: continent.sql

package db

import (
	"context"
)

const createContinent = `-- name: CreateContinent :one
INSERT INTO continents (
    continent_name,
    continent_desc) VALUES (
 $1,$2
)
RETURNING id, continent_name, continent_desc
`

type CreateContinentParams struct {
	ContinentName string `json:"continent_name"`
	ContinentDesc string `json:"continent_desc"`
}

func (q *Queries) CreateContinent(ctx context.Context, arg CreateContinentParams) (Continent, error) {
	row := q.db.QueryRow(ctx, createContinent, arg.ContinentName, arg.ContinentDesc)
	var i Continent
	err := row.Scan(&i.ID, &i.ContinentName, &i.ContinentDesc)
	return i, err
}

const deleteContinent = `-- name: DeleteContinent :exec
DELETE FROM continents
WHERE id = $1
`

func (q *Queries) DeleteContinent(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteContinent, id)
	return err
}

const getContinent0 = `-- name: GetContinent0 :one
SELECT id, continent_name, continent_desc FROM continents
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetContinent0(ctx context.Context, id int64) (Continent, error) {
	row := q.db.QueryRow(ctx, getContinent0, id)
	var i Continent
	err := row.Scan(&i.ID, &i.ContinentName, &i.ContinentDesc)
	return i, err
}

const getContinent1 = `-- name: GetContinent1 :one
SELECT id, continent_name, continent_desc FROM continents
WHERE continent_name = $1 LIMIT 1
`

func (q *Queries) GetContinent1(ctx context.Context, continentName string) (Continent, error) {
	row := q.db.QueryRow(ctx, getContinent1, continentName)
	var i Continent
	err := row.Scan(&i.ID, &i.ContinentName, &i.ContinentDesc)
	return i, err
}

const listContinents = `-- name: ListContinents :many
SELECT id, continent_name, continent_desc FROM continents
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListContinentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListContinents(ctx context.Context, arg ListContinentsParams) ([]Continent, error) {
	rows, err := q.db.Query(ctx, listContinents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Continent{}
	for rows.Next() {
		var i Continent
		if err := rows.Scan(&i.ID, &i.ContinentName, &i.ContinentDesc); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateContinent = `-- name: UpdateContinent :one
UPDATE continents
SET continent_name = $2,
continent_desc = $3
WHERE id = $1
RETURNING id, continent_name, continent_desc
`

type UpdateContinentParams struct {
	ID            int64  `json:"id"`
	ContinentName string `json:"continent_name"`
	ContinentDesc string `json:"continent_desc"`
}

func (q *Queries) UpdateContinent(ctx context.Context, arg UpdateContinentParams) (Continent, error) {
	row := q.db.QueryRow(ctx, updateContinent, arg.ID, arg.ContinentName, arg.ContinentDesc)
	var i Continent
	err := row.Scan(&i.ID, &i.ContinentName, &i.ContinentDesc)
	return i, err
}
