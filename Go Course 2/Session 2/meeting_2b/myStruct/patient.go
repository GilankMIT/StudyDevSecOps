package myStruct

import "fmt"

type Patient struct {
	Name     string
	Age      int
	Height   float64
	Weight   float64
	bmiValue float64
}

func (p Patient) GetBmiValue() float64 {
	bmiValue := p.Weight / (p.Height / 100 * p.Height / 100)
	return bmiValue
}

func (p Patient) Delete() {
	fmt.Println("DELETE")
}
