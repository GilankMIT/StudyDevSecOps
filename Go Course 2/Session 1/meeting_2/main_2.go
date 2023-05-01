package main

import (
	"fmt"
)

/*
	main() is an entry point function
	in golang
*/
func main() {
	myAge := 23 //integer variable containing a number of my age
	myAge++
	myAge--
	fmt.Println("umur saya adalah", myAge)

	instructorNames := [3]string{}
	instructorNames[2] = "Herianto"
	fmt.Println("These are the 3 names of Golang Instructor: ", instructorNames)

	indonesiaPresident := []string{"Soekarno", "Soeharto", "Habibie", "Gus Dur", "Megawati", "SBY"} //len = 6, capacity = 8
	// 6 12 24 48 dst
	fmt.Println("the current capacity of slice is ", cap(indonesiaPresident))

	indonesiaPresident = append(indonesiaPresident, "Jokowi") //len = 7
	indonesiaPresident = append(indonesiaPresident, "Gilang") //len = 8
	fmt.Println("the current capacity of slice after it reach 8 elements is ", cap(indonesiaPresident))

	indonesiaPresident = append(indonesiaPresident, "Gilang 2") //len = 9
	indonesiaPresident = append(indonesiaPresident, "Gilang 3") //len = 10
	indonesiaPresident = append(indonesiaPresident, "Gilang 4") //len = 11
	indonesiaPresident = append(indonesiaPresident, "Gilang 5") //len = 12
	fmt.Println("the current capacity of slice after it reach 12 elements is ", cap(indonesiaPresident))
	indonesiaPresident = append(indonesiaPresident, "Gilang 6") //len = 13
	fmt.Println("the current capacity of slice after it reach 13 elements is ", cap(indonesiaPresident))

	fmt.Println("These are the names of Indonesia's president: ", indonesiaPresident)

	myMap := map[string]string{}
	myMap["UserID_001"] = "Gilang"
	myMap["UserID_002"] = "Dinar"
	myMap["UserID_003"] = "Sophia"

	fmt.Println("the value of map are ", myMap)
	delete(myMap, "UserID_001")
	fmt.Println("the value of map are after deletion", myMap)
	myMap["UserID_004"] = "Irfan"
	fmt.Println("the value of map are after addition", myMap)
	myMap["UserID_001"] = "Herianto"
	fmt.Println("the value of map are after another addition", myMap)

	indonesiaPresident = []string{"Soekarno", "Soeharto", "Habibie", "Gus Dur", "Megawati", "SBY", "Jokowi"}

	//nested loop
	for i := 0; i < len(indonesiaPresident); i++ {
		fmt.Print(i+1, "the name of president : ")
		for j := 0; j < len(indonesiaPresident[i]); j++ {
			fmt.Print(string(indonesiaPresident[i][j]), "_")
		}
		fmt.Println()
	}

	for key, val := range myMap {
		fmt.Println(key, "=>", val)
	}

	additionResult := addition(15, 25)
	fmt.Println("The result will be ", additionResult)

	uselessVariable, uselessVariable2 := giveParameter("param 1", "param 2")
	fmt.Println(uselessVariable, uselessVariable2)

	voidFunction()

	addition(15, 25)

	sum(1, 2, 3, 4, 5)

	sum(1, 2, 3, 4, 5, 6, 7, 8)

	divisionResult, err := division(50, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("the result will be", divisionResult)

}

func addition(numOne int, numTwo int) int {
	result := numOne + numTwo

	return result
}

func giveParameter(param1 string, param2 string) (string, string) {
	return param1, param2
}

func voidFunction() {
	fmt.Println("Hello Gilang")
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total numebr is ", total)
}

func division(numOne, numTwo int) (int, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	return numOne / numTwo, nil
}
