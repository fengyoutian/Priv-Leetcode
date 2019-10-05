import java.util.HashMap;
import java.util.Map;
import java.util.Arrays;

class SingleNumber {

    public static void main(String[] agrs) {
        int[] numbers = new int[] { 4, 1, 2, 1, 2 };
        int result = 4;

        System.out.printf("数组: %s, 结果: %d \n", Arrays.toString(numbers), result);
        System.out.printf("HashMap 输出结果: %d \n", hashTraversing(numbers));
        System.out.printf("异或运算(XOR) 输出结果: %d \n", xor(numbers));
    }

    private static int hashTraversing(int[] nums) {
        Map<Integer, Integer> hashMap = new HashMap<>();
        for (Integer v : nums) {
            Integer count = hashMap.get(v);
            count = count == null ? 1 : ++count;
            hashMap.put(v, count);
        }

        for (Integer k : hashMap.keySet()) {
            if (hashMap.get(k) == 1) {
                return k;
            }
        }

        throw new IllegalArgumentException("Can't find solution!");
    }

    private static int xor(int[] nums) {
        int complement = 0;
        for (Integer v : nums) {
            complement ^= v;
        }

        return complement;
    }

}