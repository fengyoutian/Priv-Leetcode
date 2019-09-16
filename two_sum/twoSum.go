package main

import "fmt"

func main() {
	numbers := []int{2,7,11,15}
	result := 9

	fmt.Printf("numbers is %v, result is %d \n", numbers, result)
	fmt.Printf("violence output index is %v \n", violence(numbers, result))
	fmt.Printf("twoHashTraversing output index is %v \n", twoHashTraversing(numbers, result))
	fmt.Printf("oneHashTraversing output index is %v \n", oneHashTraversing(numbers, result))
}

func violence(numbers []int, target int) []int {
	// 暴力解法
	for i := 0; i < len(numbers) - 1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			complement := target - numbers[i]
			if numbers[j] == complement {
				return []int{i, j}
			}
		}
	}

	panic("Can not find solution!")
}

func twoHashTraversing(numbers []int, target int) []int {
	// 两遍哈希遍历
	numMap := make(map[int]int)
	for index, value := range numbers {
		numMap[value] = index
	}
	for index, value := range numbers {
		complement := target - value
		if _, ok := numMap[complement]; ok && numMap[complement] != index {
			return []int{index, numMap[complement]}
		}
	}
	panic("Can't find solution!")
}

func oneHashTraversing(numbers []int, target int) []int {
	// 一遍哈希遍历
	numMap := make(map[int]int)
	for index, value := range numbers {
		complement := target - value
		if _, ok := numMap[complement]; ok {
			return []int{numMap[complement], index}
		}
		numMap[value] = index
	}
	panic("Can't find solution!")
}