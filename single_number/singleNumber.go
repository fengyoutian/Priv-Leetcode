package main

import "fmt"

func main() {
	numbers := []int{4,1,2,1,2}
	result := 4

	fmt.Printf("数组: %v, 结果: %d \n", numbers, result)
	fmt.Printf("HashMap 输出结果: %d \n", hashTraversing(numbers))
	fmt.Printf("异或运算(XOR) 输出结果: %d \n", xor(numbers))
}

func hashTraversing(nums []int) int {
	// 看到题目，无脑选择hash表，结果一定可以算出来，空间复杂度是O(n)
	hashMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			hashMap[v] += 1
		} else {
			hashMap[v] = 1
		}
	}

	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}

	panic("Can't find solution!")
}

func xor(nums []int) int {
	// 异或特性：
	// 1. 任意数与0异或为任意数：n ^ 0 = n
	// 2. 相同数异或为0：n ^ n = 0
	// 3. 满足交换律： a ^ b ^ c = a ^ c ^ b
	complement := 0
	for _, v := range nums {
		complement ^= v
		fmt.Println(complement)
	}

	return complement
}