#!/usr/local/bin/python3
#coding=utf-8

from typing import List
import collections

class MajorityElement:
    def hash_traversing(self, nums: List[int]) -> int:
        majority = len(nums) / 2
        hash_table = {}
        for v in nums:
            if v in hash_table:
                hash_table[v] += 1
            else:
                hash_table[v] = 1
            if hash_table[v] > majority:
                return v
        raise ArithmeticError("Can't find it!")

    def hash_traversing_py(self, nums) -> int:
        '''
        leetcode写法，这里学习一下，不推荐，写算法还是少用系统函数！
        '''
        counts = collections.Counter(nums)
        return max(counts.keys(), key=counts.get)

    def array_violence(self, nums) -> int:
        majority = len(nums) / 2
        for v in nums:
            count = sum(1 for num in nums if num == v)
            if count > majority:
                return v
        raise ArithmeticError("Can't find it!")

    def sort_method(self, nums):
        '''
        同样不推荐！！！
        '''
        nums.sort()
        return nums[len(nums) // 2]

    def boyer_moore(self, nums):
        count = 0
        curr_num = None
        for v in nums:
            if count == 0:
                curr_num = v
            count += (1 if v == curr_num else -1)
        return curr_num

def main():
    numbers = [2,2,1,1,1,2,2]
    result = 2

    obj = MajorityElement()
    print("numbers: %s, result: %d" % (numbers, result))
    print("Hash表: %d" % obj.hash_traversing(numbers))
    print("Hash表(python方法): %d" % obj.hash_traversing_py(numbers))
    print("数组暴力破解: %d" % obj.array_violence(numbers))
    print("排序法: %d" % obj.sort_method(numbers))
    print("摩尔投票法: %d" % obj.boyer_moore(numbers))

if __name__ == "__main__":
    main()