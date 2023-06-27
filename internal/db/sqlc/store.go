package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

// Store 接口，装饰器
type Store interface {
	// Querier 数据库交互的全部方法接口
	Querier
	CreateUserAndProfileTx(ctx context.Context, arg CreateUserAndProfileTxParams) (CreateUserAndProfileTxResult, error)
}

// SQLStore 具体提供给外界使用的数据库操作对象
type SQLStore struct {
	// Queries 实现了和数据库交互全部方法的接口
	*Queries
	// db 数据库连接对象
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execTx 提供事务操作
// sqlc 生成的代码不提供事务操作
// 下面的代码比较死板，直接复制即可
func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateUserAndProfileTxParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Nickname       string `json:"nickname"`
	RealName       string `json:"real_name"`
	Quote          string `json:"quote"`
	Address        string `json:"address"`
}
type CreateUserAndProfileTxResult struct {
	User    User    `json:"user"`
	Profile Profile `json:"profile"`
}

// CreateUserAndProfileTx 新建用户
func (s *SQLStore) CreateUserAndProfileTx(ctx context.Context, arg CreateUserAndProfileTxParams) (CreateUserAndProfileTxResult, error) {
	var result CreateUserAndProfileTxResult
	var err error
	err = s.execTx(ctx, func(q *Queries) error {

		// 新建 user 记录
		createUserArg := CreateUserParams{
			Username:       arg.Username,
			HashedPassword: arg.HashedPassword,
			Nickname:       arg.Nickname,
		}
		result.User, err = q.CreateUser(ctx, createUserArg)
		if err != nil {
			return err
		}
		// 新建 profile 记录
		createProfileArg := CreateProfileParams{
			UserID:   result.User.ID,
			RealName: arg.RealName,
			Quote:    sql.NullString{String: arg.Quote, Valid: true},
			Address:  sql.NullString{String: arg.Address, Valid: true},
		}

		result.Profile, err = q.CreateProfile(ctx, createProfileArg)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
