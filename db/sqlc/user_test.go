package db

import (
	"context"
	"testing"
	"time"

	"github.com/naviscom/catalystx2/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomName(8),
		HashedPassword: hashedPassword,
		FullName:       util.RandomName(8),
		Email:          util.RandomEmail(),
		Role:           util.RandomName(8),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.PasswordCreatedAt)
	require.Equal(t, arg.Role, user.Role)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser0(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser0(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.PasswordCreatedAt, user2.PasswordCreatedAt, time.Second)
	require.Equal(t, user1.Role, user2.Role)
}

func TestGetUser3(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser3(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.PasswordCreatedAt, user2.PasswordCreatedAt, time.Second)
	require.Equal(t, user1.Role, user2.Role)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)

	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := UpdateUserParams{
		Username:          user1.Username,
		HashedPassword:    hashedPassword,
		FullName:          util.RandomName(8),
		Email:             util.RandomName(8),
		PasswordChangedAt: time.Now().UTC(),
		Role:              util.RandomName(8),
	}
	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.Email, user2.Email)
	require.WithinDuration(t, arg.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.Equal(t, arg.Role, user2.Role)

}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)
	user2, err := testStore.GetUser0(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)

}
