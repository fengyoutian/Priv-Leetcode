import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class TwoSum {

    public static void main(String[] args) {
        int numbers[] = new int[] { 2, 7, 11, 15 };
        int result = 9;

        System.out.printf("numbers is %s, result is %d\n", Arrays.toString(numbers), result);
        System.out.printf("violence output index is %s\n", Arrays.toString(violence(numbers, result)));
        System.out.printf("twoHashTraversing output index is %s\n", Arrays.toString(twoHashTraversing(numbers, result)));
        System.out.printf("oneHashTraversing output index is %s\n", Arrays.toString(oneHashTraversing(numbers, result)));
    }

    /**
     * 暴力方式
     * 
     * @param nums
     * @param target
     * @return
     */
    private static int[] violence(int[] nums, int target) {
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                int complement = target - nums[i];
                if (complement == nums[j]) {
                    return new int[] { i, j };
                }
            }
        }
        throw new IllegalArgumentException("Can't find solution!");
    }

    /**
     * 两遍哈希
     * 
     * @param nums
     * @param target
     * @return
     */
    private static int[] twoHashTraversing(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }

        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement) && map.get(complement) != i) {
                return new int[] { i, map.get(complement) };
            }
        }

        throw new IllegalArgumentException("Can't find solution!");
    }

    /**
     * 一遍哈希
     * 
     * @param nums
     * @param target
     * @return
     */
    private static int[] oneHashTraversing(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] { map.get(complement), i };
            }
            map.put(nums[i], i);
        }

        throw new IllegalArgumentException("Can't find solution!");
    }

}