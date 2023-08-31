package db

import (
	"context"
	"simple-gobank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer{
	arg := CreateTransferParams{
		FromAccountID: account1.ID, 
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t, createRandomAccount(t), createRandomAccount(t))
}

func TestGetTransfer(t *testing.T){
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, account1, account2)

	receivedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, receivedTransfer)

	require.Equal(t, transfer.ID, receivedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, receivedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, receivedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, receivedTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, receivedTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T){
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 5,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
