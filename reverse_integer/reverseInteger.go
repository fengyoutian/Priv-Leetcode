package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("INT_MAX: %d, INT_MIN: %d\n", INT_MAX, INT_MIN)
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", 1234, intConv2Str(1234))
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", -1234, intConv2Str(-1234))
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", 0, intConv2Str(0))
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", INT_MAX, intConv2Str(INT_MAX))
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", INT_MIN, intConv2Str(INT_MIN))
	fmt.Printf("整数转成字符串, 输入: %d, 输出: %d\n", 1534236469, intConv2Str(1534236469))
	fmt.Println("===============================")
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", 1234, reverse(1234))
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", -1234, reverse(-1234))
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", 0, reverse(0))
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", INT_MAX, reverse(INT_MAX))
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", INT_MIN, reverse(INT_MIN))
	fmt.Printf("数学计算, 输入: %d, 输出: %d\n", 1534236469, reverse(1534236469))
}

// 定义int最大最小值
const (
	// INT_MAX int = int(^uint32(0) >> 1)
	INT_MAX int = 1 << 31 - 1
	INT_MIN int = ^INT_MAX
)

func intConv2Str(x int) int {
	isNegative := x < 0
	str := strconv.Itoa(x)
	if isNegative {
		str = str[1:] // 先去掉负号
	}
	revStr := ""
	for i := range str {
		revStr = str[i: i + 1] + revStr
	}
	// 先取出个位数, 判断是否溢出后再放回
	pop, _ := strconv.Atoi(revStr[len(revStr) - 1:])
	x, _ = strconv.Atoi(revStr[: len(revStr) - 1])
	// 判断是否溢出
	if !isNegative && (x < (INT_MAX / 10) || (x == (INT_MAX / 10) && pop <= (INT_MAX % 10))) {
		return x * 10 + pop
	} else if isNegative && (x < (INT_MAX / 10) || (x == (INT_MAX / 10) && pop <= (INT_MAX % 10 + 1))) {
		return -(x * 10 + pop)
	}

	return 0
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		// 溢出判断
		if (rev > (INT_MAX / 10) || (rev == (INT_MAX / 10) && pop > (INT_MAX % 10))) || 
			(rev < (INT_MIN / 10) || (rev == (INT_MIN / 10) && pop < (INT_MIN % 10))) {
			return 0
		}

		rev = rev * 10 + pop
	}

	return rev
}