// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: state.sql

package db

import (
	"context"
)

const createState = `-- name: CreateState :one
INSERT INTO states (
    state_name,
    state_desc,
    country_id,
    area_id) VALUES (
 $1,$2,$3,$4
)
RETURNING id, state_name, state_desc, country_id, area_id
`

type CreateStateParams struct {
	StateName string `json:"state_name"`
	StateDesc string `json:"state_desc"`
	CountryID int64  `json:"country_id"`
	AreaID    int64  `json:"area_id"`
}

func (q *Queries) CreateState(ctx context.Context, arg CreateStateParams) (State, error) {
	row := q.db.QueryRow(ctx, createState,
		arg.StateName,
		arg.StateDesc,
		arg.CountryID,
		arg.AreaID,
	)
	var i State
	err := row.Scan(
		&i.ID,
		&i.StateName,
		&i.StateDesc,
		&i.CountryID,
		&i.AreaID,
	)
	return i, err
}

const deleteState = `-- name: DeleteState :exec
DELETE FROM states
WHERE id = $1
`

func (q *Queries) DeleteState(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteState, id)
	return err
}

const getState0 = `-- name: GetState0 :one
SELECT id, state_name, state_desc, country_id, area_id FROM states
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetState0(ctx context.Context, id int64) (State, error) {
	row := q.db.QueryRow(ctx, getState0, id)
	var i State
	err := row.Scan(
		&i.ID,
		&i.StateName,
		&i.StateDesc,
		&i.CountryID,
		&i.AreaID,
	)
	return i, err
}

const getState1 = `-- name: GetState1 :one
SELECT id, state_name, state_desc, country_id, area_id FROM states
WHERE state_name = $1 LIMIT 1
`

func (q *Queries) GetState1(ctx context.Context, stateName string) (State, error) {
	row := q.db.QueryRow(ctx, getState1, stateName)
	var i State
	err := row.Scan(
		&i.ID,
		&i.StateName,
		&i.StateDesc,
		&i.CountryID,
		&i.AreaID,
	)
	return i, err
}

const listStates = `-- name: ListStates :many
SELECT id, state_name, state_desc, country_id, area_id FROM states
WHERE country_id = $3 OR area_id = $4
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStatesParams struct {
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
	CountryID int64 `json:"country_id"`
	AreaID    int64 `json:"area_id"`
}

func (q *Queries) ListStates(ctx context.Context, arg ListStatesParams) ([]State, error) {
	rows, err := q.db.Query(ctx, listStates,
		arg.Limit,
		arg.Offset,
		arg.CountryID,
		arg.AreaID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []State{}
	for rows.Next() {
		var i State
		if err := rows.Scan(
			&i.ID,
			&i.StateName,
			&i.StateDesc,
			&i.CountryID,
			&i.AreaID,
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

const updateState = `-- name: UpdateState :one
UPDATE states
SET state_name = $2,
state_desc = $3,
country_id = $4,
area_id = $5
WHERE id = $1
RETURNING id, state_name, state_desc, country_id, area_id
`

type UpdateStateParams struct {
	ID        int64  `json:"id"`
	StateName string `json:"state_name"`
	StateDesc string `json:"state_desc"`
	CountryID int64  `json:"country_id"`
	AreaID    int64  `json:"area_id"`
}

func (q *Queries) UpdateState(ctx context.Context, arg UpdateStateParams) (State, error) {
	row := q.db.QueryRow(ctx, updateState,
		arg.ID,
		arg.StateName,
		arg.StateDesc,
		arg.CountryID,
		arg.AreaID,
	)
	var i State
	err := row.Scan(
		&i.ID,
		&i.StateName,
		&i.StateDesc,
		&i.CountryID,
		&i.AreaID,
	)
	return i, err
}
