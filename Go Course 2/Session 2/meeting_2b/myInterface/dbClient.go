package myInterface

import (
	"meeting_2/myStruct"
)

type DBClient interface {
	Insert(patient myStruct.Patient) error
}
