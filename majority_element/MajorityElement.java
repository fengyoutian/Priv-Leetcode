import java.util.HashMap;
import java.util.Map;
import java.util.Arrays;

class MajorityElement {

    public static void main(String[] args) {
        int[] numbers = new int[] { 2, 2, 1, 1, 1, 2, 2 };
        int result = 2;

        System.out.printf("numbers: %s, result: %d\n", Arrays.toString(numbers), result);
        System.out.printf("Hash表: %d\n", hashTraversing(numbers));
        System.out.printf("数组暴力解法: %d\n", arrayViolence(numbers));
        System.out.printf("排序法: %d\n", sortMethod(numbers));
        System.out.printf("摩尔投票法: %d\n", boyerMoore(numbers));
    }

    private static int hashTraversing(int[] nums) {
        Map<Integer, Integer> hashMap = new HashMap<>();
        int majority = nums.length / 2;
        for (Integer v : nums) {
            Integer count = hashMap.get(v);
            count = (count == null ? 1 : ++count);
            hashMap.put(v, count);

            if (count > majority) {
                return v;
            }
        }
        throw new IllegalArgumentException("Can't find it!");
    }

    private static int arrayViolence(int[] nums) {
        int majority = nums.length / 2;
        for (int v : nums) {
            int count = 0;
            for (int num : nums) {
                if (v == num) {
                    count++;
                }
                if (count > majority) {
                    return v;
                }
            }
        }
        throw new IllegalArgumentException("Can't find it!");
    }

    private static int sortMethod(int[] nums) {
        Arrays.sort(nums);
        return nums[nums.length / 2];
    }

    private static int boyerMoore(int[] nums) {
        int count = 0;
        Integer currNum = null;
        for (int v : nums) {
            if (count == 0) {
                currNum = v;
            }
            count += currNum == v ? 1 : -1;
        }
        return currNum;
    }

}