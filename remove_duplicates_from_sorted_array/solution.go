package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := []int{0, 1, 2, 3, 4}

	fmt.Printf("nums: %v, result: %v, len: %d\n", nums, result, len(result))
	len := doublePointer(nums)
	fmt.Printf("双指针 result: %d, nums: %v\n", len, nums[:len])
}

func doublePointer(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1 // 返回的是len，需要+1
}
