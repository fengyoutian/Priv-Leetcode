#!/usr/local/bin/python3
#coding=utf-8

from typing import List

class Solution:
    def violence(self, nums: List[int], target: int) -> List[int]:
        """
        :暴力解法 (letcode超时)
        """
        i = 0
        while i < len(nums) - 1:
            j = i + 1
            while j < len(nums):
                complement = target - nums[i]
                if complement == nums[j]:
                    return [i, j]
                j += 1
            i += 1
        raise ValueError("Can't find solution!")

    def violence_2(self, nums: List[int], target: int) -> List[int]:
        """
        :暴力解法2
        """
        for i in range(0, len(nums) - 1):
            for j in range(i + 1, len(nums)):
                complement = target - nums[i]
                if complement == nums[j]:
                    return [i, j]

        raise ValueError("Can't find solution!")


    def two_hash_traversing(self, nums: List[int], target: int) -> List[int]:
        """
        :两遍哈希表
        """
        hashmap = {}
        for index in range(len(nums)):
            hashmap[nums[index]] = index
        
        for index, value in enumerate(nums):
            complement = target - value
            if complement in hashmap and hashmap[complement] != index:
                return [index, hashmap[complement]]

        raise ValueError("Can't find solution!")

    def one_hash_traversing(self, nums: List[int], target: int) -> List[int]:
        """
        :一遍哈希表
        """
        hashmap = {}
        for index, value in enumerate(nums):
            complement = target - value
            if complement in hashmap:
                return [hashmap[complement], index]
            hashmap[value] = index

        raise ValueError("Can't find solution!")

def main():
    numbers = [2,7,11,15]
    result = 9

    solution = Solution()
    print("numbers is %s, result is %d" % (numbers, result))
    print("violence output index is %s" % (solution.violence(numbers, result)))
    print("violence_2 output index is %s" % (solution.violence_2(numbers, result)))
    print("two_hash_traversing output index is %s" % (solution.two_hash_traversing(numbers, result)))
    print("one_hash_traversing output index is %s" % (solution.one_hash_traversing(numbers, result)))

if __name__ == "__main__":
    main()