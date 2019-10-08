package main

import "fmt"

func main() {
	matrix := [][]int {
		{1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 20
	target_1 := 5
	
	fmt.Printf("matrix: %v\n", matrix)
	fmt.Printf("排除法: %t, target: %d\n", searchMatrix(matrix, target), target)
	fmt.Printf("排除法: %t, target: %d\n", searchMatrix(matrix, target_1), target_1)
	// 这个结果在LeetCode上执行出错，数组越界，不知道咋回事。。。
	matrix_1 := [][]int {{-1, -1}}
	fmt.Printf("排除法: %t, matrix: %v , target: %d\n", searchMatrix(matrix_1, 0), matrix_1, 0)
}

func searchMatrix(matrix [][]int, target int) bool {
    row_max := len(matrix)
    if row_max == 0 {
        return false
    }
    
    col_max := len(matrix[0])
    if col_max == 0 {
        return false
    }
    
    row := row_max - 1
    col := 0
    for ; row >= 0 && col < row_max; {
        if target == matrix[row][col] {
            return true
        } else if target > matrix[row][col] {
            col++
        } else {
            row--
        }
    }
    
    return false
}