class SingleNumber {
    hashTraversing(nums) {
        const hash_map = {}
        for (const v of nums) {
            if (hash_map[v]) {
                hash_map[v]++
            } else {
                hash_map[v] = 1
            }
        }

        for (const k of Object.keys(hash_map)) {
            if (hash_map[k] == 1) {
                return k
            }
        }

        throw Error("Can't find solution!")
    }

    xor(nums) {
        let complement = 0
        for (const v of nums) {
            complement ^= v
        }
        return complement
    }

}

const numbers = [4, 1, 2, 1, 2]
const result = 4

const obj = new SingleNumber()
console.log("数组: %s, 结果: %d", numbers, result)
console.log("HashMap 输出结果: %d", obj.hashTraversing(numbers))
console.log("异或运算(XOR) 输出结果: %d", obj.xor(numbers))