#!/usr/bin/python3
#coding=utf-8

from typing import List
import time
import random

class Solution:
    def bubbleSort(self, nums: List[int]) -> List[int]:
        for i in range(len(nums) - 1):
            for j in range(i, len(nums)):
                if nums[i] > nums[j]:
                    nums[i], nums[j] = nums[j], nums[i]
        return nums

    def insertSort(self, nums: List[int]) -> List[int]:
        for i in range(1, len(nums)):
            j = i
            while j - 1 >= 0:
                if nums[j - 1] > nums[j]:
                    nums[j - 1], nums[j] = nums[j], nums[j - 1]
                j -= 1
        return nums

    def selectionSort(self, nums: List[int]) -> List[int]:
        for i in range(len(nums)):
            min = i
            for j in range(i, len(nums)):
                if nums[min] > nums[j]:
                    min = j
            nums[min], nums[i] = nums[i], nums[min]
        return nums

    def shellSort(self, nums: List[int]) -> List[int]:
        gap = len(nums) // 2
        while gap >= 1:
            for i in range(gap, len(nums)):
                j = i
                while j - 1 >= 0:
                    if nums[j - 1] > nums[j]:
                        nums[j - 1], nums[j] = nums[j], nums[j - 1]
                    j -= gap
            gap -= 1
        return nums

    def quickSort(self, nums: List[int]) -> List[int]:
        '''
            挖坑填数
        '''
        def quick(low, high):
            if low >= high: return
            mid = (low + high) // 2
            key = nums[mid]
            nums[low], nums[mid] = nums[mid], nums[low]

            l, r = low, high
            while l < r:
                while l < r and nums[r] >= key:
                # while l < r:
                    # if nums[r] < key: break
                    r -= 1 
                nums[l] = nums[r]
                while l < r and nums[l] <= key:
                # while l < r:
                    # if nums[l] > key: break
                    l += 1
                nums[r] = nums[l]
            nums[l] = key
            quick(low, l - 1)
            quick(l + 1, high)
        quick(0, len(nums) - 1)
        return nums

    def quickSort2(self, nums: List[int]) -> List[int]:
        '''
            双指针交换
        '''
        def quick(low, high):
            if low >= high: return
            mid = (low + high) // 2
            key = nums[mid]
            nums[low], nums[mid] = nums[mid], nums[low]
            l, r = low, high
            while l < r:
                while l < r and nums[r] >= key: r -= 1
                while l < r and nums[l] <= key: l += 1
                if l >= r: break
                nums[l], nums[r] = nums[r], nums[l]
            nums[l], nums[low] = nums[low], nums[l]
            quick(low, l - 1)
            quick(l + 1, high)
        quick(0, len(nums) - 1)
        return nums

    def quickSort3(self, nums: List[int]) -> List[int]:
        '''
            单指针交换
        '''
        def quick(low, high):
            if low >= high: return
            mid = (low + high) // 2
            key = nums[mid]
            nums[low], nums[mid] = nums[mid], nums[low]
            l = low
            for j in range(low + 1, high + 1):
                if nums[j] < key:
                    l += 1
                    nums[l], nums[j] = nums[j], nums[l]
            nums[low], nums[l] = nums[l], nums[low]
            quick(low, l - 1)
            quick(l + 1, high)
        quick(0, len(nums) - 1)
        return nums

    def mergeSort(self, nums: List[int]) -> List[int]:
        def merge(low, high):
            if low >= high: return
            mid = (low + high) // 2
            merge(low, mid)
            merge(mid + 1, high)
            self.merge(nums, low, mid, high)
        merge(0, len(nums) - 1)
        return nums

    def merge(self, nums: List[int], low, mid, high: int):
        # 临时数组
        tmp = [0] * (high - low + 1)
        # 左、右、临时数组坐标
        l, r, i = low, mid + 1, 0
        # 按升序归并到数组
        while l <= mid and r <= high:
            if nums[l] <= nums[r]:
                tmp[i] = nums[l]
                l += 1
            else:
                tmp[i] = nums[r]
                r += 1
            i += 1
        # 右边复制完成，左边剩余，直接复制左边剩余数据到临时数组
        while l <= mid:
            tmp[i] = nums[l]
            l += 1
            i += 1
        # 左边复制完成，右边剩余，直接复制右边剩余数据到临时数组
        while r <= high:
            tmp[i] = nums[r]
            r += 1
            i += 1
        # 将临时数组复制回原数组 len(tmp) == i
        for i in range(len(tmp)):
            nums[low + i] = tmp[i]

    def mergeSort2(self, nums: List[int]) -> List[int]:
        '''
        循环实现归并排序
        '''
        low, mid, high, length = 0, 0, 0, len(nums) - 1
        # 分割长度，从1开始
        split = 1
        while split <= length:
            i = 0 
            while i + split <= length:
                low, mid, high = i, i + split - 1, i + 2*split - 1
                # 防止高位超过数组长度
                if high > length: high = length
                self.merge(nums, low, mid, high)
                i += 2 * split
            # 增加切割单位
            split *= 2
        return nums

def main():
    nums = [0, 11, 3, 24, 21, 2, 10, 20, 18, 1, 
            26, 16, 4, 28, 6, 15, 5, 8, 19, 9, 
            12, 29, 27, 13, 14, 17, 7, 22, 23, 25]
    print("nums: %s\n" % nums)

    obj = Solution()
    start = time.time()
    print("Bubble Sort result: %s, time: %fs" % (obj.bubbleSort(nums[:]), time.time() - start))
    start = time.time()
    print("Insert Sort result: %s, time: %fs" % (obj.insertSort(nums[:]), time.time() - start))
    start = time.time()
    print("Selection Sort result: %s, time: %fs" % (obj.selectionSort(nums[:]), time.time() - start))
    start = time.time()
    print("Shell Sort result: %s, time: %fs" % (obj.shellSort(nums[:]), time.time() - start))
    start = time.time()
    print("Quick Sort result: %s, time: %fs" % (obj.quickSort(nums[:]), time.time() - start))
    start = time.time()
    print("Quick Sort 2 result: %s, time: %fs" % (obj.quickSort2(nums[:]), time.time() - start))
    start = time.time()
    print("Quick Sort 3 result: %s, time: %fs" % (obj.quickSort3(nums[:]), time.time() - start))
    start = time.time()
    print("Merge Sort result: %s, time: %fs" % (obj.mergeSort(nums[:]), time.time() - start))
    start = time.time()
    print("Merge Sort 2 result: %s, time: %fs" % (obj.mergeSort2(nums[:]), time.time() - start))

if __name__ == "__main__":
    main()