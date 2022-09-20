package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBank(t *testing.T) {
	arg := CreateBankParams{
		BankName:      "GTB",
		AccountNumber: "0213050718",
		ShopID:        1,
	}
	bank, err := testQueries.CreateBank(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, bank)
	require.Equal(t, arg.BankName, bank.BankName)
	require.Equal(t, arg.AccountNumber, bank.AccountNumber)
	require.Equal(t, arg.ShopID, bank.ShopID)
	require.NotZero(t, bank.ID)
	require.NotZero(t, bank.CreatedAt)
}
