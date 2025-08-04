package main

import "fmt"

func main() {
	sumIndex := towSum([]int{2, 3, 45, 7}, 9)
	fmt.Printf("sumIndex: %+v\n", sumIndex)
	sum2 := towSum2([]int{2, 3, 5, 7}, 7)
	fmt.Printf("sum2: %+v\n", sum2)
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

// 每次都是从当前数向前，找补数。这样会避免找到这个元素本身

func towSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num

		// 判断补数，是否在map中
		idx, found := numMap[complement]
		if found {
			return []int{idx, i}
		}
		// 将数字集合放入map中，k是 值，v是index
		numMap[num] = i
	}
	return nil
}
