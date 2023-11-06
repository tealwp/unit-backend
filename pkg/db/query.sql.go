// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package db

import (
	"context"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
  name
) VALUES (
  $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreateEvent(ctx context.Context, name string) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent, name)
	var i Event
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT id, name, created_at FROM events
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int64) (Event, error) {
	row := q.db.QueryRow(ctx, getEvent, id)
	var i Event
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT id, name, created_at FROM events
ORDER BY name
`

func (q *Queries) ListEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEvent = `-- name: UpdateEvent :exec
UPDATE events
  set name = $2
WHERE id = $1
`

type UpdateEventParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) error {
	_, err := q.db.Exec(ctx, updateEvent, arg.ID, arg.Name)
	return err
}
