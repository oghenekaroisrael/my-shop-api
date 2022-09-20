package db

import (
	"context"
	"testing"
	"time"

	"github.com/oghenekaroisrael/myshopapi/utils"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashPwd, err := utils.HashPassword(utils.RandomPassword())
	arg := CreateUserParams{
		FirstName: utils.RandomName(),
		LastName:  utils.RandomName(),
		Email:     utils.RandomEmail(),
		Password:  hashPwd,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUserByEmailAndPassword(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmailAndPassword(context.Background(), GetUserByEmailAndPasswordParams{user1.Email, user1.Password})
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserDetail(t *testing.T) {
	user1 := createRandomUser(t)
	args := UpdateUserDetailParams{
		ID:        user1.ID,
		FirstName: utils.RandomName(),
		LastName:  utils.RandomName(),
	}
	user2, err := testQueries.UpdateUserDetail(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, args.FirstName, user2.FirstName)
	require.Equal(t, args.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserPassword(t *testing.T) {
	user1 := createRandomUser(t)
	args := UpdateUserPasswordParams{
		ID:       user1.ID,
		Email:    user1.Email,
		Password: utils.RandomPassword(),
	}
	user2, err := testQueries.UpdateUserPassword(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, args.Password, user2.Password)
}

func TestUpdateUserVerificationStatus(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.UpdateUserVerificationStatus(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, true, user2.IsVerified)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	args := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, users, 5)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
