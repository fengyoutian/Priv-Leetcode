import java.util.Arrays;

class Search2DMatrixII {
    public boolean searchMatrix(int[][] matrix, int target) {
        int row_max = matrix.length;
        if (row_max == 0) {
            return false;
        }
        int col_max = matrix[0].length;
        if (col_max == 0) {
            return false;
        }

        int row = row_max - 1;
        int col = 0;
        while (row >= 0 && col < col_max) {
            if (target == matrix[row][col]) {
                return true;
            } else if (target > matrix[row][col]) {
                col = col + 1;
            } else {
                row = row - 1;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        int[][] matrix = new int[][] {
            {1,   4,  7, 11, 15},
            {2,   5,  8, 12, 19},
            {3,   6,  9, 16, 22},
            {10, 13, 14, 17, 24},
            {18, 21, 23, 26, 30}
        };
        int target = 20;
        int target_1 = 5;

        Search2DMatrixII solution = new Search2DMatrixII();
        System.out.printf("排除法: %b, target: %d\n", solution.searchMatrix(matrix, target), target);
        System.out.printf("排除法: %b, target: %d\n", solution.searchMatrix(matrix, target_1), target_1);
    }
}