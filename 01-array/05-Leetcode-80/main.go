package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	k := 2 // [0, k) 表示Remove Duplicates后的数组区间

	for i := 2; i < length; i++ {
		if nums[i] > nums[k-2] {
			nums[k] = nums[i]
			k += 1
		}
	}

	return k
}

func main() {
}
