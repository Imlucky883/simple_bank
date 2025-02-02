package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := CreateAccountParams{
		Owner:    "postgres",
		Balance:  1000,
		Currency: "USD",
	}

	createdAccount, _ := testQueries.CreateAccount(context.Background(), account)

	entry := CreateEntryParams{
		AccountID: createdAccount.ID,
		Amount:    500,
	}

	createdEntry, err := testQueries.CreateEntry(context.Background(), entry)
	require.NoError(t, err)
	require.NotEmpty(t, createdEntry)
	require.Equal(t, entry.AccountID, createdEntry.AccountID)
	require.Equal(t, entry.Amount, createdEntry.Amount)
}
