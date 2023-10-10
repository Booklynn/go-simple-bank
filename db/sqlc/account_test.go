package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Booklynn/simplebank/ulti"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    ulti.RandomOwner(),
		Balance:  ulti.RandomMoney(),
		Currency: ulti.RandomCurrency(),
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

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	expectedAccount := createRandomAccount(t)

	actualAccount, err := testQueries.GetAccount(context.Background(), expectedAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	require.Equal(t, expectedAccount.ID, actualAccount.ID)
	require.Equal(t, expectedAccount.Owner, actualAccount.Owner)
	require.Equal(t, expectedAccount.Balance, actualAccount.Balance)
	require.Equal(t, expectedAccount.Currency, actualAccount.Currency)
	require.WithinDuration(t, expectedAccount.CreatedAt, actualAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: ulti.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	errDel := testQueries.DeleteAccount(context.Background(), createdAccount.ID)

	detletedAccount, errGet := testQueries.GetAccount(context.Background(), createdAccount.ID)

	require.NoError(t, errDel)
	require.Error(t, errGet)
	require.EqualError(t, errGet, sql.ErrNoRows.Error())
	require.Empty(t, detletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	createdAccounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, createdAccounts, 5)

	for _, account := range createdAccounts {
		require.NotEmpty(t, account)
	}
}
