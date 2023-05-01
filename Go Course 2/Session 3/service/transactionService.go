package service

import(
	"testing_example/dbClient"
	"testing_example/model"
)

type TransactionService struct{
	DBClient dbClient.DBClient
}

func (t TransactionService) CalculateTotalAmountOfOrder() int64{
	myOrdersAny := t.DBClient.Query("SELECT * FROM order")

	myOrders := myOrdersAny.([]model.Order)

	var totalAmt int64 = 0

	for _, order := range myOrders{
		totalAmt += order.Amount
	}

	return totalAmt
}