package db

import (
	"context"
	"testing"
	"time"

	"github.com/Booklynn/simplebank/ulti"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    ulti.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	expectedEntry := createRandomEntry(t, account)

	actaulEntry, err := testQueries.GetEntry(context.Background(), expectedEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actaulEntry)

	require.Equal(t, expectedEntry.ID, actaulEntry.ID)
	require.Equal(t, expectedEntry.AccountID, actaulEntry.AccountID)
	require.Equal(t, expectedEntry.Amount, actaulEntry.Amount)
	require.WithinDuration(t, expectedEntry.CreatedAt, actaulEntry.CreatedAt, time.Second)
}

func TestListEntires(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := listEntiresParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entires, err := testQueries.listEntires(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entires, 5)

	for _, entry := range entires {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
