
function violence(nums, target) {
    // 暴力解法
    for (i = 0; i < nums.length - 1; i++) {
        for (j = i + 1; j < nums.length; j++) {
            let complement = target - nums[i]
            if (nums[j] == complement) {
                return [i, j]
            }
        }
    }
    throw Error("Can't find solution!")
}

function twoHashTraversing(nums, target) {
    // 两遍哈希
    let hashMap = {}
    nums.forEach((value, index) => {
        hashMap[value] = index
    });

    for (index in nums) {
        let complement = target - nums[index]
        if (hashMap[complement] != void 0 && hashMap[complement] != index) {
            return [Number(index), hashMap[complement]]
        }
    }
    throw Error("Can't find solution!")
}

function oneHashTraversing(nums, target) {
    // 一遍哈希
    let hashMap = {}
    for (index in nums) {
        let complement = target - nums[index]
        if (hashMap[complement] != void 0) {
            return [Number(hashMap[complement]), Number(index)]
        }
        hashMap[nums[index]] = index
    }
    throw Error("Can't find solution!")
}

let numbers = [2,7,11,15]
let result = 9

console.log("numbers is %s, result is %d", numbers, result)
console.log("violence output index is %s", violence(numbers, result))
console.log("twoHashTraversing output index is %s", twoHashTraversing(numbers, result))
console.log("oneHashTraversing output index is %s", oneHashTraversing(numbers, result))