package db

import (
	"context"
	"github.com/Keekuun/go_simple_bank/util"
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
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)

	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestGetEntryList(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := GetEntryListParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	list, err := testQueries.GetEntryList(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, list, 5)

	for _, entry := range list {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
