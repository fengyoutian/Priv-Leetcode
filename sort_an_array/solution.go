package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{
		0, 11, 3, 24, 21, 2, 10, 20, 18, 1,
		26, 16, 4, 28, 6, 15, 5, 8, 19, 9,
		12, 29, 27, 13, 14, 17, 7, 22, 23, 25,
	}

	start := time.Now()
	fmt.Printf("Bubble Sort result: %v, time: %s\n", bubbleSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Insert Sort result: %v, time: %s\n", insertSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Selection Sort result: %v, time: %s\n", selectionSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Shell Sort result: %v, time: %s\n", shellSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Quick Sort result: %v, time: %s\n", quickSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Merge Sort result: %v, time: %s\n", mergeSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Merge Sort 2 result: %v, time: %s\n", mergeSort2(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
	fmt.Printf("Heap Sort result: %v, time: %s\n", heapSort(append([]int{}, nums...)), time.Since(start))
	start = time.Now()
}

// 冒泡排序
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				// 使用 加减法 交换元素
				nums[i] = nums[i] + nums[j]
				nums[j] = nums[i] - nums[j]
				nums[i] = nums[i] - nums[j]
			}
		}
	}
	return nums
}

// 插入排序
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j-1 >= 0 && nums[j-1] > nums[j]; j-- {
			// 使用 异或运算 交换元素
			nums[j-1] ^= nums[j]
			nums[j] ^= nums[j-1]
			nums[j-1] ^= nums[j]
		}
	}

	return nums
}

// 选择排序
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		// 直接使用Golang特性交换元素
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

// 希尔排序
func shellSort(nums []int) []int {
	length := len(nums)
	// 增量直接从数组一半开始
	gap := length / 2
	// 增量递减，直到1
	for gap >= 1 {
		// 插入排序，按增量（gap）拆分数组
		for i := gap; i < length; i++ {
			for j := i; j-1 >= 0 && nums[j-1] > nums[j]; j -= gap {
				// 使用临时变量交换元素
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			}
		}
		// 增量递减
		gap /= 2
	}
	return nums
}

// 快速排序
func quickSort(nums []int) []int {
	return quick(nums, 0, len(nums)-1)
}

func quick(nums []int, low, high int) []int {
	if low >= high {
		return nums
	}
	mid := (low + high) / 2
	key := nums[mid]
	nums[low], nums[mid] = nums[mid], nums[low]

	l, r := low, high
	for l < r {
		// 右指针向左移动，找到大于key的坐标，并和左指针交换
		for l < r && nums[r] >= key {
			r--
		}
		nums[l] = nums[r]
		// 左指针向右移动，找到小于key的坐标，并和右指针交换
		for l < r && nums[l] <= key {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = key          // 循环结束后 key 放回左指针才算交换完成
	quick(nums, low, l-1)  // 继续排序左半部分
	quick(nums, l+1, high) // 继续排序右半部分
	return nums
}

// 归并排序 (递归方式)
func mergeSort(nums []int) []int {
	return mergeSortFunc(nums, 0, len(nums)-1)
}

func mergeSortFunc(nums []int, low, high int) []int {
	if low >= high {
		return nums
	}
	middle := low + (high-low)/2        // 二分
	mergeSortFunc(nums, low, middle)    // 拆分左半边
	mergeSortFunc(nums, middle+1, high) // 拆分右半边
	merge(nums, low, middle, high)      // 归并左右两边
	return nums
}

func merge(nums []int, low, middle, high int) {
	length := high - low + 1        // 归并后数组长度
	tempNums := make([]int, length) // 临时合并数组
	// 左、右、临时数组起始坐标
	left, right, index := low, middle+1, 0
	// 按升序归并到新数组中
	for ; left <= middle && right <= high; index++ {
		if nums[left] <= nums[right] {
			tempNums[index] = nums[left]
			left++
		} else {
			tempNums[index] = nums[right]
			right++
		}
	}
	// 右边数组复制完成，左边剩余，则依次复制到临时数组中
	for ; left <= middle; index++ {
		tempNums[index] = nums[left]
		left++
	}
	// 左边数组复制完成，右边剩余，则依次复制到临时数组中
	for ; right <= high; index++ {
		tempNums[index] = nums[right]
		right++
	}
	// 将归并后的数组复制到原数组中
	for i := 0; i < length; i++ {
		nums[low+i] = tempNums[i]
	}
}

// 归并排序（非递归实现）
func mergeSort2(nums []int) []int {
	low, middle, high, length := 0, 0, 0, len(nums)-1
	// 待归并数组长度：1 2 4 8 ...
	split := 1 // 从最小分割单位 1 开始
	for split <= length {
		// 按分割单位遍历数组并合并
		for i := 0; i+split <= length; i += split * 2 {
			low = i
			// middle 主要是在合并时找到右半边数组的起始下标
			middle = i + split - 1
			high = i + 2*split - 1
			// 防止超过数组长度
			if high > length {
				high = length
			}
			// 归并两个有序的子数组
			merge(nums, low, middle, high)
		}
		// 增加切分单位
		split *= 2
	}

	return nums
}

func heapSort(nums []int) []int {
	n := len(nums)
	// 构建堆，一开始可将数组看做无序的堆
	// 将下标为 n/2 开始到 0 的元素下沉到合适的位置
	// 因为 n/2 后面的元素都是椰子节点，无需下沉
	for p := n / 2; p >= 0; p-- {
		adjust(nums, p, n)
	}

	// 下沉排序
	// 堆的根节点永远是最大值，所以只需将最大值和最后一位的元素交换即可
	// 然后再维护一个除原最大节点意外的 n-1 的堆，再将新堆的根节点放在倒数第二的位置，如此反复
	for i := n - 1; i > 0; i-- {
		// 将 nums[0] 与最大元素 a[i] 交换，并修复堆
		nums[0], nums[i] = nums[i], nums[0]
		// 下沉排序，重新构建
		adjust(nums, 0, i)
	}

	return nums
}

func adjust(nums []int, p, n int) {
	// 是否存在左孩子节点
	for 2*p+1 < n {
		// 左孩节点下标
		left := 2*p + 1
		right := 2*p + 2
		children := left
		// right < n 说明存在右孩子，判断将根节点下沉到左还是右
		// 如果左孩小于右孩，那么下沉到右子树的根，并且下次从右子树开始判断是否还要下沉
		if right < n && nums[left] < nums[right] {
			children = right
		}

		// 如果根节点不小于它的子节点，表示这个子树根节点最大
		if nums[p] >= nums[children] {
			break // 不用下沉，直接跳出
		}

		// 否则将根节点下沉为它的左子树或右子树的根，也就是将较大的值上调
		nums[p], nums[children] = nums[children], nums[p]

		// 继续从左子树或右子树开始，判断根节点是否还要下沉
		p = children
	}
}
