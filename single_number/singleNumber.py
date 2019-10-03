#!/usr/local/bin/python3
#coding=utf-8

from typing import List

class SingleNumber:
    def hashTraversing(self, nums: List[int]) -> int:
        hash_map = {}
        for i in nums:
            if i in hash_map:
                hash_map[i] += 1
            else:
                hash_map[i] = 1

        for k, v in hash_map.items():
            if v == 1:
                return k
        
        raise ValueError("Can't find it")

    def xor(self, nums) -> int:
        complement = 0
        for v in nums:
            complement ^= v
        return complement

def main():
    numbers = [4,1,2,1,2]
    result = 4

    obj = SingleNumber()
    print("数组: %s, 结果: %d" % (numbers, result))
    print("HashMap 输出结果: %d" % obj.hashTraversing(numbers))
    print("异或运算(XOR) 输出结果: %d \n" % obj.xor(numbers))

if __name__ == "__main__":
    main()