package tests

import (
	"context"
	"github.com/ngenohkevin/bitscrow/config"
	db "github.com/ngenohkevin/bitscrow/db/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransaction(t *testing.T) db.EscrowTransaction {
	arg := db.CreateEscrowTransactionsParams{
		ID:             config.GenerateRandomID(),
		BuyerID:        config.GenerateRandomID(),
		SellerID:       config.GenerateRandomID(),
		BitcoinAddress: config.RandomBitcoinAddress(),
		Amount:         config.RandomAmount(),
		Status:         config.RandomStatus(),
	}

	transaction, err := testStore.CreateEscrowTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.ID, transaction.ID)
	require.Equal(t, arg.BuyerID, transaction.BuyerID)
	require.Equal(t, arg.SellerID, transaction.SellerID)
	require.Equal(t, arg.BitcoinAddress, transaction.BitcoinAddress)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.Status, transaction.Status)
	require.NotZero(t, transaction.CreatedAt)

	return transaction
}

func TestCreateEscrowTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetEscrowTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	transaction2, err := testStore.GetEscrowTransactions(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.BuyerID, transaction2.BuyerID)
	require.Equal(t, transaction1.SellerID, transaction2.SellerID)
	require.Equal(t, transaction1.BitcoinAddress, transaction2.BitcoinAddress)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Status, transaction2.Status)
	//require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)

}
