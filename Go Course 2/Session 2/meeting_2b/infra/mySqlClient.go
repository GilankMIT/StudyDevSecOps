package infra

import (
	"fmt"
	"meeting_2/myStruct"
)

type MySQLDBClient struct {
}

func (db MySQLDBClient) Insert(patient myStruct.Patient) error {
	fmt.Println("Insert data ", patient.Name)
	return nil
}
