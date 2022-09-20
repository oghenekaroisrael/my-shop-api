// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: token.sql

package db

import (
	"context"
	"database/sql"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  user_id,
  access_token,
  refresh_token
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, access_token, refresh_token
`

type CreateTokenParams struct {
	UserID       int32          `json:"user_id"`
	AccessToken  sql.NullString `json:"access_token"`
	RefreshToken sql.NullString `json:"refresh_token"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken, arg.UserID, arg.AccessToken, arg.RefreshToken)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
	)
	return i, err
}

const getUserAccessToken = `-- name: GetUserAccessToken :one
SELECT user_id, access_token FROM tokens
WHERE user_id = $1
LIMIT 1
`

type GetUserAccessTokenRow struct {
	UserID      int32          `json:"user_id"`
	AccessToken sql.NullString `json:"access_token"`
}

func (q *Queries) GetUserAccessToken(ctx context.Context, userID int32) (GetUserAccessTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getUserAccessToken, userID)
	var i GetUserAccessTokenRow
	err := row.Scan(&i.UserID, &i.AccessToken)
	return i, err
}

const getUserRefresToken = `-- name: GetUserRefresToken :one
SELECT user_id, refresh_token FROM tokens
WHERE user_id = $1
LIMIT 1
`

type GetUserRefresTokenRow struct {
	UserID       int32          `json:"user_id"`
	RefreshToken sql.NullString `json:"refresh_token"`
}

func (q *Queries) GetUserRefresToken(ctx context.Context, userID int32) (GetUserRefresTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getUserRefresToken, userID)
	var i GetUserRefresTokenRow
	err := row.Scan(&i.UserID, &i.RefreshToken)
	return i, err
}
