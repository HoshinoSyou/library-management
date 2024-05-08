package main

import (
	"library-management/cmd"
	"library-management/dao"
)

func main() {
	defer dao.DB.Close()
	cmd.Entrance()
}

//func maximumCount(nums []int) int {
//	neg := Search(nums, 0)
//	pos := len(nums) - Search(nums, 1)
//	return int(math.Max(float64(neg), float64(pos)))
//}
//
//func Search(nums []int, val int) int {
//	left := 0
//	right := len(nums)
//	for left < right {
//		i := (left + right) / 2
//		if nums[i] >= val {
//			right = i
//		} else {
//			left = i + 1
//		}
//	}
//	return left
//}
