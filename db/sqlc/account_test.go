package db

import (
	"context"
	"database/sql"
	"simple-gobank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
  user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner: user.Username,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}	

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	savedAccount := createRandomAccount(t)
	receivedAccount, err := testQueries.GetAccount(context.Background(), savedAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, receivedAccount)

	require.Equal(t, savedAccount.ID, receivedAccount.ID)
	require.Equal(t, savedAccount.Owner, receivedAccount.Owner)
	require.Equal(t, savedAccount.Balance, receivedAccount.Balance)
	require.Equal(t, savedAccount.Currency, receivedAccount.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, receivedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	savedAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID: savedAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, savedAccount.ID, updatedAccount.ID)
	require.Equal(t, savedAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, savedAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}	

func TestDeleteAccount(t *testing.T){
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	someAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, someAccount)
}

func TestListAccounts(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAccount(t)	
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}