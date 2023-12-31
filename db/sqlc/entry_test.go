package db

import (
	"context"
	"course/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
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
	account1 := createRandomAccount(t)
	createRandomEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := createRandomEntry(t, account1)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account1 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account1)
	}

	arg := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	list, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, entry := range list {
		require.NotEmpty(t, entry)
	}
}
