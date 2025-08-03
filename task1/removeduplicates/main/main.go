package main

import "fmt"

func main() {
	count := removeDuplicates([]int{1, 2, 2, 4, 5})
	fmt.Printf("count: %+v\n", count)
}

func removeDuplicates(nums []int) int {
label:
	for i := 0; i < len(nums)-1; i++ {
		currentNum := nums[i]
		nextNum := nums[i+1]
		if currentNum == nextNum {
			nums = append(nums[:i], nums[i+1:]...)
			goto label
		}
	}
	return len(nums)
}
