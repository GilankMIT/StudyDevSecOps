package service

import (
	"testing"
	"testing_example/dbClient"
	"testing_example/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTotalAmountFromOrder(t *testing.T){
	sampleOrders := []model.Order{
		{
			ID: "Order 1",
			Amount: 1000,
		},
		{
			ID: "Order 2",
			Amount: 1500,
		},
	}

	dbClientMock := new(dbClient.MockDBClient)
	dbClientMock.On("Query", mock.Anything).Return(sampleOrders)

	TransactionService := TransactionService{
		DBClient: dbClientMock,
	}

	totalAmount := TransactionService.CalculateTotalAmountOfOrder()

	assert.Equal(t, int64(2500), totalAmount, "error not equal")
}