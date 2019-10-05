package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 2, 1, 1, 1, 2, 2}
	result := 2

	fmt.Printf("numbers: %v, result: %d\n", numbers, result)
	fmt.Printf("Hash表: %d\n", oneHashTraversing(numbers))
	fmt.Printf("Hash表，两个for循环: %d\n", twoHashTraversing(numbers))
	fmt.Printf("数组暴力破解: %d\n", arrayViolence(numbers))
	fmt.Printf("排序法: %d\n", sortMethod(numbers))
	fmt.Printf("摩尔投票法: %d\n", boyerMoore(numbers))
}

func oneHashTraversing(nums []int) int {
	hash_table := make(map[int]int)
	majority := len(nums) / 2
	for _, v := range nums {
		hash_table[v]++
		if hash_table[v] > majority {
			return v
		}
	}

	panic("Can't find it!")
}

func twoHashTraversing(nums []int) int {
	hash_table := make(map[int]int)
	majority := len(nums) / 2
	for _, v := range nums {
		hash_table[v]++
	}

	for k, v := range hash_table {
		if v > majority {
			return k
		}
	}

	panic("Can't find it!")
}

func arrayViolence(nums []int) int {
	len := len(nums)
	majority := len / 2
	for i := 0; i < len; i++ {
		count := 0
		for j := 0; j < len; j++ {
			if nums[i] == nums[j] {
				count++
			}

			if count > majority {
				return nums[i]
			}
		}
	}

	panic("Can't find it!")
}

func sortMethod(nums []int) int {
	// 不推荐使用，既然是算法，还是自己写吧！
	sort.Ints(nums)
	return nums[len(nums) / 2]
}

func boyerMoore(nums []int) int {
	count := 0
	var curr_num int
	for _, v := range nums {
		if count == 0 {
			curr_num = v
		}
		if v == curr_num {
			count++
		} else {
			count--
		}
	}
	return curr_num
}