// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: sitetype.sql

package db

import (
	"context"
)

const createSitetype = `-- name: CreateSitetype :one
INSERT INTO sitetypes (
    type_name,
    type_desc) VALUES (
 $1,$2
)
RETURNING id, type_name, type_desc
`

type CreateSitetypeParams struct {
	TypeName string `json:"type_name"`
	TypeDesc string `json:"type_desc"`
}

func (q *Queries) CreateSitetype(ctx context.Context, arg CreateSitetypeParams) (Sitetype, error) {
	row := q.db.QueryRow(ctx, createSitetype, arg.TypeName, arg.TypeDesc)
	var i Sitetype
	err := row.Scan(&i.ID, &i.TypeName, &i.TypeDesc)
	return i, err
}

const deleteSitetype = `-- name: DeleteSitetype :exec
DELETE FROM sitetypes
WHERE id = $1
`

func (q *Queries) DeleteSitetype(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSitetype, id)
	return err
}

const getSitetype0 = `-- name: GetSitetype0 :one
SELECT id, type_name, type_desc FROM sitetypes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSitetype0(ctx context.Context, id int64) (Sitetype, error) {
	row := q.db.QueryRow(ctx, getSitetype0, id)
	var i Sitetype
	err := row.Scan(&i.ID, &i.TypeName, &i.TypeDesc)
	return i, err
}

const getSitetype1 = `-- name: GetSitetype1 :one
SELECT id, type_name, type_desc FROM sitetypes
WHERE type_name = $1 LIMIT 1
`

func (q *Queries) GetSitetype1(ctx context.Context, typeName string) (Sitetype, error) {
	row := q.db.QueryRow(ctx, getSitetype1, typeName)
	var i Sitetype
	err := row.Scan(&i.ID, &i.TypeName, &i.TypeDesc)
	return i, err
}

const listSitetypes = `-- name: ListSitetypes :many
SELECT id, type_name, type_desc FROM sitetypes
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSitetypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSitetypes(ctx context.Context, arg ListSitetypesParams) ([]Sitetype, error) {
	rows, err := q.db.Query(ctx, listSitetypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sitetype{}
	for rows.Next() {
		var i Sitetype
		if err := rows.Scan(&i.ID, &i.TypeName, &i.TypeDesc); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSitetype = `-- name: UpdateSitetype :one
UPDATE sitetypes
SET type_name = $2,
type_desc = $3
WHERE id = $1
RETURNING id, type_name, type_desc
`

type UpdateSitetypeParams struct {
	ID       int64  `json:"id"`
	TypeName string `json:"type_name"`
	TypeDesc string `json:"type_desc"`
}

func (q *Queries) UpdateSitetype(ctx context.Context, arg UpdateSitetypeParams) (Sitetype, error) {
	row := q.db.QueryRow(ctx, updateSitetype, arg.ID, arg.TypeName, arg.TypeDesc)
	var i Sitetype
	err := row.Scan(&i.ID, &i.TypeName, &i.TypeDesc)
	return i, err
}
