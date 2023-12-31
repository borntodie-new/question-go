// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username,
                   nickname,
                   hashed_password)
VALUES ($1, $2, $3) RETURNING id, username, hashed_password, nickname, avatar, status, is_super, created_at, updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Nickname, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
UPDATE users
SET status = $2
WHERE id = $1 RETURNING id, username, hashed_password, nickname, avatar, status, is_super, created_at, updated_at
`

type DeleteUserParams struct {
	ID     int64        `json:"id"`
	Status sql.NullBool `json:"status"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, arg.ID, arg.Status)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const queryUsers = `-- name: QueryUsers :many
SELECT u.id, u.username, u.hashed_password, u.nickname, u.avatar, u.status, u.is_super, u.created_at, u.updated_at,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id LIMIT $1
OFFSET $2
`

type QueryUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type QueryUsersRow struct {
	ID             int64          `json:"id"`
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	Nickname       string         `json:"nickname"`
	Avatar         sql.NullString `json:"avatar"`
	Status         sql.NullBool   `json:"status"`
	IsSuper        sql.NullBool   `json:"is_super"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	RealName       sql.NullString `json:"real_name"`
	Gender         sql.NullInt32  `json:"gender"`
	Quote          sql.NullString `json:"quote"`
	Address        sql.NullString `json:"address"`
}

func (q *Queries) QueryUsers(ctx context.Context, arg QueryUsersParams) ([]QueryUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, queryUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QueryUsersRow{}
	for rows.Next() {
		var i QueryUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.Nickname,
			&i.Avatar,
			&i.Status,
			&i.IsSuper,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.RealName,
			&i.Gender,
			&i.Quote,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveUserByID = `-- name: RetrieveUserByID :one
SELECT u.id, u.username, u.hashed_password, u.nickname, u.avatar, u.status, u.is_super, u.created_at, u.updated_at,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id
WHERE u.id = $1 LIMIT 1
`

type RetrieveUserByIDRow struct {
	ID             int64          `json:"id"`
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	Nickname       string         `json:"nickname"`
	Avatar         sql.NullString `json:"avatar"`
	Status         sql.NullBool   `json:"status"`
	IsSuper        sql.NullBool   `json:"is_super"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	RealName       sql.NullString `json:"real_name"`
	Gender         sql.NullInt32  `json:"gender"`
	Quote          sql.NullString `json:"quote"`
	Address        sql.NullString `json:"address"`
}

func (q *Queries) RetrieveUserByID(ctx context.Context, id int64) (RetrieveUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, retrieveUserByID, id)
	var i RetrieveUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RealName,
		&i.Gender,
		&i.Quote,
		&i.Address,
	)
	return i, err
}

const retrieveUserByUsername = `-- name: RetrieveUserByUsername :one
SELECT u.id, u.username, u.hashed_password, u.nickname, u.avatar, u.status, u.is_super, u.created_at, u.updated_at,
       p.real_name,
       p.gender,
       p.quote,
       p.address
FROM users AS u
         LEFT JOIN profiles AS p ON u.id = p.user_id
WHERE u.username = $1 LIMIT 1
`

type RetrieveUserByUsernameRow struct {
	ID             int64          `json:"id"`
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	Nickname       string         `json:"nickname"`
	Avatar         sql.NullString `json:"avatar"`
	Status         sql.NullBool   `json:"status"`
	IsSuper        sql.NullBool   `json:"is_super"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	RealName       sql.NullString `json:"real_name"`
	Gender         sql.NullInt32  `json:"gender"`
	Quote          sql.NullString `json:"quote"`
	Address        sql.NullString `json:"address"`
}

func (q *Queries) RetrieveUserByUsername(ctx context.Context, username string) (RetrieveUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, retrieveUserByUsername, username)
	var i RetrieveUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RealName,
		&i.Gender,
		&i.Quote,
		&i.Address,
	)
	return i, err
}

const updateAdmin = `-- name: UpdateAdmin :one
UPDATE users
SET is_super = $2
WHERE id = $1 RETURNING id, username, hashed_password, nickname, avatar, status, is_super, created_at, updated_at
`

type UpdateAdminParams struct {
	ID      int64        `json:"id"`
	IsSuper sql.NullBool `json:"is_super"`
}

func (q *Queries) UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateAdmin, arg.ID, arg.IsSuper)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users
SET hashed_password = $2
WHERE id = $1 RETURNING id, username, hashed_password, nickname, avatar, status, is_super, created_at, updated_at
`

type UpdatePasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Nickname,
		&i.Avatar,
		&i.Status,
		&i.IsSuper,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
