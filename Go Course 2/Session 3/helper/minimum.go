package helper

import "errors"

func GetMinimum(nums []int) (int, error){
	
	if nums == nil{
		return 0, errors.New("nums cannot be null")
	}

	if len(nums) < 1{
		return 0, nil
	}

	smallestNumber := nums[0]
	for i := 1; i < len(nums); i++{
		if nums[i] < smallestNumber{
			smallestNumber = nums[i]
		}
	}

	return smallestNumber, nil
}