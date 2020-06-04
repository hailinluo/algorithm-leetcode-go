package main

import "fmt"

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		// NOTE：使用 (l+r)/2 可能产生整型溢出
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 5, 8, 10, 34, 78, 99}
	ret := binarySearch(arr, 34)
	fmt.Println(ret)
}
