package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 7, 11, 15}
	var target int = 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(nums []int, target int) []int {
	var sum []int

	seenNums := make(map[int]int)
	for index, num := range nums {
		seenIndex, numExists := seenNums[target-num]
		if numExists {
			sum = append(sum, seenIndex)
			sum = append(sum, index)
			return sum
		}
		seenNums[num] = index
	}
	return sum
}
