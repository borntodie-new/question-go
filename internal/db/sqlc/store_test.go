package sqlc

import (
	"context"
	"github.com/borntodie-new/question-go/internal/db/helper"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateUserAndProfileTx(t *testing.T) {
	arg := CreateUserAndProfileTxParams{
		Username:       helper.RandomOwner(),
		HashedPassword: helper.RandomString(32),
		Nickname:       helper.RandomOwner(),
		RealName:       helper.RandomOwner(),
		Quote:          helper.RandomString(24),
		Address:        helper.RandomString(32),
	}
	createUserAndProfileTxResult, err := testStore.CreateUserAndProfileTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createUserAndProfileTxResult)

	require.Equal(t, arg.Username, createUserAndProfileTxResult.User.Username)
	require.Equal(t, arg.HashedPassword, createUserAndProfileTxResult.User.HashedPassword)
	require.Equal(t, arg.Nickname, createUserAndProfileTxResult.User.Nickname)
	require.Equal(t, arg.RealName, createUserAndProfileTxResult.Profile.RealName)
	require.Equal(t, arg.Quote, createUserAndProfileTxResult.Profile.Quote.String)
	require.Equal(t, arg.Address, createUserAndProfileTxResult.Profile.Address.String)

	require.False(t, createUserAndProfileTxResult.User.IsSuper.Bool)
	require.True(t, createUserAndProfileTxResult.User.Status.Bool)
	require.WithinDuration(t, time.Now(), createUserAndProfileTxResult.User.CreatedAt, time.Second)
	require.WithinDuration(t, time.Now(), createUserAndProfileTxResult.User.UpdatedAt, time.Second)

	require.Equal(t, createUserAndProfileTxResult.User.ID, createUserAndProfileTxResult.Profile.UserID)
	require.WithinDuration(t, createUserAndProfileTxResult.User.CreatedAt, createUserAndProfileTxResult.Profile.CreatedAt, time.Second)
	require.WithinDuration(t, createUserAndProfileTxResult.User.CreatedAt, createUserAndProfileTxResult.Profile.UpdatedAt, time.Second)

}
