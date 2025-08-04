package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {3, 6}, {8, 9}, {10, 13}}
	m := merge(intervals)
	fmt.Printf("m: %+v\n", m)
}

func merge(intervals [][]int) [][]int {
	// 基础判断
	if len(intervals) <= 1 {
		return intervals
	}

	// 	排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	//循环对比merged 的最后一个区间，和 intervals 的每个区间
	for i := 1; i < len(intervals); i++ {
		// 合并结果的最后一个区间
		last := merged[len(merged)-1]
		// intervals的每个区间
		interval := intervals[i]
		// 对比intervals的 每个区间开始值大于 合并结果的最后一个区间的最大值，就把这个区间加到合并区间
		if interval[0] > last[1] {
			merged = append(merged, interval)
		} else {
			// 当intervals的区间开始节点，小于等于 合并区间的最大值，说明要进行合并区间修正，
			// 将合并区间的右侧值进行修正。如果intervals的结束节点大于 merged区间的时候。
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		}
	}

	return merged
}
