package main

import (
	"fmt"
	"meeting_2/infra"
	"meeting_2/myInterface"
	"meeting_2/myStruct"
	"time"
)

func main() {
	patientZero := myStruct.Patient{
		Name:   "Gilang Prambudi",
		Height: 175,
		Weight: 60,
		Age:    23,
	}


	patientZero.Delete()

	fmt.Println("Patient ", patientZero.Name, " BMI is ", patientZero.GetBmiValue())

	dBClient := infra.MySQLDBClient{}

	go Insert(patientZero, dBClient) //takes 3 second
	go Insert(patientZero, dBClient) //takes 3 second
	go Insert(patientZero, dBClient) //takes 3 second
	go Insert(patientZero, dBClient) //takes 3 second
	go Insert(patientZero, dBClient) //takes 3 second
	go Insert(patientZero, dBClient) //takes 3 second

	fmt.Println("sending email to user for new registration")

	time.Sleep(5 * time.Second)
}

func Insert(p myStruct.Patient, dbClient myInterface.DBClient) {
	time.Sleep(3 * time.Second)
	dbClient.Insert(p)
}
