package sqlc

import (
	"context"
	"database/sql"
	"github.com/borntodie-new/question-go/internal/db/helper"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func randomCreateProfile(t *testing.T) Profile {
	user := randomCreateUser(t)
	arg := CreateProfileParams{
		UserID:   user.ID,
		RealName: helper.RandomOwner(),
		Quote:    sql.NullString{String: helper.RandomString(32), Valid: true},
		Address:  sql.NullString{String: helper.RandomString(32), Valid: true},
	}
	profile, err := testStore.CreateProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, arg.RealName, profile.RealName)
	require.Equal(t, arg.Quote, profile.Quote)
	require.Equal(t, arg.Address, profile.Address)
	require.Equal(t, arg.UserID, profile.UserID)

	require.WithinDuration(t, user.CreatedAt, profile.CreatedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, profile.UpdatedAt, time.Second)

	return profile
}

func TestCreateProfile(t *testing.T) {
	randomCreateProfile(t)
}
