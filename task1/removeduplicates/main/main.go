package main

import "fmt"

func main() {
	count := removeDuplicates([]int{1, 2, 2, 4, 5})
	fmt.Printf("count: %+v\n", count)

	duplicates1 := removeDuplicates1([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Printf("duplicates1: %+v\n", duplicates1)
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

func removeDuplicates1(nums []int) int {
	// 使用快慢两个指针
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
