package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinimum_normalFlow(t *testing.T) {
	sampleInput := []int{1, 2, 3, -10, -200, 2321, 1021213}
	expectedOutput := -200

	//calling the function that want to be tested
	actualResult, err := GetMinimum(sampleInput)


	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, actualResult, expectedOutput, "actual result and expected result not match")
}

func TestGetMinimum_zeroLenArray(t *testing.T){
	sampleInput := []int{}
	expectedOutput := 0

	//calling the function that want to be tested
	actualResult, err := GetMinimum(sampleInput)

	assert.Nil(t, err, "error must be nil")
	assert.Equal(t, actualResult, expectedOutput, "actual result and expected result not match")
}

func TestGetMinimum_nilArray(t *testing.T){
	expectedOutput := 0

	//calling the function that want to be tested
	actualResult, err := GetMinimum(nil)

	assert.NotNil(t, err, "error must not be nil")
	assert.Equal(t, actualResult, expectedOutput, "actual result and expected result not match")
}