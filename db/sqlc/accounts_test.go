package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	account := CreateAccountParams{
		Owner:    "postgres",
		Balance:  1000,
		Currency: "USD",
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), account)
	require.NoError(t, err)
	require.NotEmpty(t, createdAccount)
	require.Equal(t, account.Owner, createdAccount.Owner)
	require.Equal(t, account.Balance, createdAccount.Balance)
	require.Equal(t, account.Currency, createdAccount.Currency)
}

func TestGetAccount(t *testing.T) {
	account := CreateAccountParams{
		Owner:    "postgres",
		Balance:  1000,
		Currency: "USD",
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), account)
	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)
}
