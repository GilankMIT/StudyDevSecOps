package main

import (
	"fmt"
	"reflect"
)

/*
	main() is an entry point function that will be
	executed in Go first

	author: Gilang Prambudi
*/
func main() {

	//this is a variable with value 10
	//this is the next line of comment
	a := 10

	var b int = 50

	c := a + b

	fmt.Println(c)
	fmt.Printf("%s is my name\n\n", "gilang")

	//variables:

	//1. boolean
	var isUserActive bool = false
	fmt.Println("is user active? ", isUserActive)

	//2.
	var myAge int = 128
	fmt.Println("my age is ", myAge, " years old")

	//2b.
	var totalCirculatingSupply int64 = 1_000_000_000_000_000 //1000000000000000
	fmt.Println("total circulating supply ", totalCirculatingSupply)

	//3.
	var phi float32 = 3.14000001
	fmt.Printf("the value for phi is %.05f\n", phi)

	//4.
	var myName string = "Gifari"
	fmt.Println("My name is ", myName)

	//logical operator
	var aNumber int = 10
	var bNumber int = 20
	fmt.Println("this expression is ", aNumber > bNumber)
	fmt.Println("this expression is ", aNumber < bNumber)
	fmt.Println("this expression is ", aNumber >= bNumber)
	fmt.Println("this expression is ", aNumber <= bNumber)
	fmt.Println("this expression is ", aNumber == bNumber)
	fmt.Println("this expression is ", aNumber != bNumber)

	myFloatingNumber1 := 10.5 //float
	fmt.Println("the value will be", myFloatingNumber1/2)

	myUnknownDataType := 1_000_000_000.50
	fmt.Println("the data type is ", reflect.TypeOf(myUnknownDataType))



	minCustomerAge := 17
	maxCustomerAge := 50
	myCurrentAge := 75

	if myCurrentAge >= minCustomerAge && myCurrentAge <= maxCustomerAge{
		
		fmt.Println("I can buy this product")

	}else if myCurrentAge == 22 {

		fmt.Println("I can still buy this product")

	}else{

		fmt.Println("No, I cannot buy this product")

	}




	//my list
	myListOfNumbers := []int{1, 2, 3, 4, 5} //total is = 15
	theSumOfListOfNumbers := 0
	
	//loop
	// for index := len(myListOfNumbers) - 1; index >= 0; index--{
	// 	theSumOfListOfNumbers += myListOfNumbers[index]
	// }

	index := 0
	for index < len(myListOfNumbers){
		theSumOfListOfNumbers += myListOfNumbers[index]
		index++
	}

	fmt.Println("the total value of myListOfNUmber is", theSumOfListOfNumbers)

}
