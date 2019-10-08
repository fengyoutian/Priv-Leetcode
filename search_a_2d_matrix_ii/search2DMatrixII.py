#!/usr/local/bin/python3
#coding=utf-8

from typing import List

class Solution:
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        row_max = len(matrix)
        if row_max == 0:
            return False
        col_max = len(matrix[0])
        if col_max == 0:
            return False

        row = row_max - 1
        col = 0
        while row >= 0 and col < col_max:
            if target == matrix[row][col]:
                return True
            elif target > matrix[row][col]:
                col = col + 1
            else:
                row = row - 1
        return False

def main():
    matrix = [
        [1,   4,  7, 11, 15],
        [2,   5,  8, 12, 19],
        [3,   6,  9, 16, 22],
        [10, 13, 14, 17, 24],
        [18, 21, 23, 26, 30]
    ]
    target = 20
    target_1 = 5

    solution = Solution()
    print("matrix: {}".format(matrix))
    print("排除法: {}, target: {}".format(solution.searchMatrix(matrix, target), target))
    print("排除法: {}, target: {}".format(solution.searchMatrix(matrix, target_1), target_1))

if __name__ == "__main__":
    main()