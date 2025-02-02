package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := CreateAccountParams{
		Owner:    "sender",
		Balance:  2000,
		Currency: "USD",
	}

	account2 := CreateAccountParams{
		Owner:    "receiver",
		Balance:  1000,
		Currency: "USD",
	}

	createdAccount1, _ := testQueries.CreateAccount(context.Background(), account1)
	createdAccount2, _ := testQueries.CreateAccount(context.Background(), account2)

	transfer := CreateTransferParams{
		FromAccountID: createdAccount1.ID,
		ToAccountID:   createdAccount2.ID,
		Amount:        300,
	}

	createdTransfer, err := testQueries.CreateTransfer(context.Background(), transfer)
	require.NoError(t, err)
	require.NotEmpty(t, createdTransfer)
	require.Equal(t, transfer.FromAccountID, createdTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, createdTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, createdTransfer.Amount)
}
