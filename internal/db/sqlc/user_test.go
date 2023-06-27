package sqlc

import (
	"context"
	"database/sql"
	"github.com/borntodie-new/backend-master-class/util"
	"github.com/borntodie-new/question-go/internal/db/helper"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func randomCreateUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		Nickname:       util.RandomOwner(),
		HashedPassword: util.RandomString(32),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Nickname, user.Nickname)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	randomCreateUser(t)
}

func TestRetrieveUserByUsername(t *testing.T) {
	user1 := randomCreateUser(t)

	user2, err := testStore.RetrieveUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.IsSuper, user2.IsSuper)
	require.Equal(t, user1.Status, user2.Status)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)

	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

func TestRetrieveUserByID(t *testing.T) {
	user1 := randomCreateUser(t)

	user2, err := testStore.RetrieveUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.IsSuper, user2.IsSuper)
	require.Equal(t, user1.Status, user2.Status)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)

	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

func TestQueryUsers(t *testing.T) {
	users1 := make([]User, 0)
	for i := 0; i < 5; i++ {
		user := randomCreateUser(t)
		users1 = append(users1, user)
	}

	arg := QueryUsersParams{
		Limit:  5,
		Offset: 0,
	}
	users2, err := testStore.QueryUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users2)

	for _, user := range users2 {
		require.NotEmpty(t, user)
	}

}

func TestDeleteUser(t *testing.T) {
	user1 := randomCreateUser(t)

	arg := DeleteUserParams{
		ID:     user1.ID,
		Status: sql.NullBool{Bool: true, Valid: true},
	}
	user2, err := testStore.DeleteUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.IsSuper, user2.IsSuper)
	require.Equal(t, true, user2.Status.Bool)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)

	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdatePassword(t *testing.T) {
	user1 := randomCreateUser(t)

	arg := UpdatePasswordParams{
		ID:             user1.ID,
		HashedPassword: helper.RandomString(32),
	}
	user2, err := testStore.UpdatePassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.IsSuper, user2.IsSuper)
	require.Equal(t, user1.Status, user2.Status)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.NotEqual(t, user1.HashedPassword, user2.HashedPassword)

	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)

	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateAdmin(t *testing.T) {
	user1 := randomCreateUser(t)

	arg := UpdateAdminParams{
		ID:      user1.ID,
		IsSuper: sql.NullBool{Bool: true, Valid: true},
	}
	user2, err := testStore.UpdateAdmin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Nickname, user2.Nickname)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, true, user2.IsSuper.Bool)
	require.Equal(t, user2.Status, user2.Status)
	require.Equal(t, user1.Avatar, user2.Avatar)

	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)

	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
