package main

import "fmt"

func main() {
	sumIndex := towSum([]int{2, 3, 45, 7}, 9)
	fmt.Printf("sumIndex: %+v\n", sumIndex)
}

func towSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
