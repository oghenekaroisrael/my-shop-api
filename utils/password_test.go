package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hash1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)

	err = CheckPassword(password, hash1)
	require.NoError(t, err)

	wrongpassword := RandomString(6)
	err = CheckPassword(wrongpassword, hash1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hash2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash2)
	require.NotEqual(t, hash1, hash2)
}
