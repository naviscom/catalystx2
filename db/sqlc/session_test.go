package db

import (
	"context"
	"testing"
	"time"

	"github.com/naviscom/catalystx2/util"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T, user User) Session {
	arg := CreateSessionParams{
		Username:     user.USERNAME,
		RefreshToken: util.RandomName(8),
		UserAgent:    util.RandomName(8),
		ClientIp:     util.RandomName(8),
		ExpiresAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
	}
	session, err := testStore.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	require.Equal(t, arg.ID, session.ID)
	require.Equal(t, arg.Username, session.Username)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.ClientIp, session.ClientIp)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)
	require.WithinDuration(t, arg.ExpiresAt, session.ExpiresAt, time.Second)
	require.WithinDuration(t, arg.CreatedAt, session.CreatedAt, time.Second)
	return session
}

func TestCreateSession(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
}

func TestGetSession0(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	session2, err := testStore.GetSession0(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.Username, session2.Username)
	require.Equal(t, session1.RefreshToken, session2.RefreshToken)
	require.Equal(t, session1.UserAgent, session2.UserAgent)
	require.Equal(t, session1.ClientIp, session2.ClientIp)
	require.Equal(t, session1.IsBlocked, session2.IsBlocked)
	require.WithinDuration(t, session1.ExpiresAt, session2.ExpiresAt, time.Second)
	require.WithinDuration(t, session1.CreatedAt, session2.CreatedAt, time.Second)
}

func TestListSessions(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomSession(t, user)

	}
	arg := ListSessionsParams{
		Username: user.USERNAME,
		Limit:    5,
		Offset:   5,
	}
	sessions, err := testStore.ListSessions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, sessions, 5)

	for _, session := range sessions {
		require.NotEmpty(t, session)
		require.True(t, arg.Username == session.Username)
	}
}

func TestUpdateSession(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	arg := UpdateSessionParams{
		Username:     user.USERNAME,
		RefreshToken: util.RandomName(8),
		UserAgent:    util.RandomName(8),
		ClientIp:     util.RandomName(8),
		ExpiresAt:    time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
	}
	session2, err := testStore.UpdateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.Username, session2.Username)
	require.Equal(t, arg.RefreshToken, session2.RefreshToken)
	require.Equal(t, arg.UserAgent, session2.UserAgent)
	require.Equal(t, arg.ClientIp, session2.ClientIp)
	require.Equal(t, arg.IsBlocked, session2.IsBlocked)
	require.WithinDuration(t, arg.ExpiresAt, session2.ExpiresAt, time.Second)
	require.WithinDuration(t, arg.CreatedAt, session2.CreatedAt, time.Second)

}

func TestDeleteSession(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	err := testStore.DeleteSession(context.Background(), session1.ID)
	require.NoError(t, err)
	session2, err := testStore.GetSession0(context.Background(), session1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, session2)

}
