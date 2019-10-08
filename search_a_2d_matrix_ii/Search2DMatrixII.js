/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    const row_max = matrix.length;
    if (row_max == 0) {
        return false;
    }
    const col_max = matrix[0].length;
    if (col_max == 0) {
        return false;
    }

    let row = row_max - 1;
    let col = 0;
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
};

const matrix = [
    [1,   4,  7, 11, 15],
    [2,   5,  8, 12, 19],
    [3,   6,  9, 16, 22],
    [10, 13, 14, 17, 24],
    [18, 21, 23, 26, 30]
]
const target = 20
const target_1 = 5

console.log("matrix", matrix)
console.log("排除法: %d, target: %d", searchMatrix(matrix, target), target)
console.log("排除法: %d, target: %d", searchMatrix(matrix, target_1), target_1)