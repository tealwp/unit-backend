// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
}
