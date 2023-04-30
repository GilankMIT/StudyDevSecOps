package infra

import (
	"fmt"
	"meeting_2/myStruct"
)

type MongoDBClient struct {
}

func (db MongoDBClient) Insert(patient myStruct.Patient) error {
	fmt.Println("Insert data to MongoDB", patient.Name)
	return nil
}
