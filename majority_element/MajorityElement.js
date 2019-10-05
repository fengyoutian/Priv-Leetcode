
class MajorityElement {

    constructor() {
    }

    hashTraversing(nums) {
        const hashMap = {}
        const majority = nums.length / 2
        for (const v of nums) {
            hashMap[v] = hashMap[v] == void 0 ? 1 : hashMap[v] + 1
            if (hashMap[v] > majority) {
                return v
            }
        }
        throw Error("Can't find it!")
    }

    arrayViolence(nums) {
        const majority = nums.length / 2
        for (const v of nums) {
            let count = 0
            for (const num of nums) {
                if (v == num) {
                    count++
                }
                if (count > majority) {
                    return v
                }
            }
        }
        throw Error("Can't find it!")
    }

    sortMethod(nums) {
        nums.sort()
        return nums[Math.trunc(nums.length / 2)]
    }

    boyerMoore(nums) {
        const majority = nums.length / 2
        let count = 0
        let currNum = undefined
        for (const v of nums) {
            if (count == 0) {
                currNum = v
            }
            count += v == currNum ? 1 : -1
        }
        return currNum
    }

}

const numbers = [2,2,1,1,1,2,2]
const result = 2

const obj = new MajorityElement()
console.log("numbers: %s, result: %d", numbers, result)
console.log("Hash表: %d", obj.hashTraversing(numbers))
console.log("数组暴力破解: %d", obj.arrayViolence(numbers))
console.log("排序法: %d", obj.sortMethod(numbers))
console.log("摩尔投票法: %d", obj.boyerMoore(numbers))