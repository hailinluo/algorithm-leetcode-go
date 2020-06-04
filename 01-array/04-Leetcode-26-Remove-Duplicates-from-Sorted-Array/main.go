package main

func main() {
	
}

func removeDuplicates(nums []int) int {
	var k int // [0, k] 表示不重复的元素
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k] {
			if i != k+1 {
				nums[k+1] = nums[i]
			}
			k += 1
		}
	}

	return k + 1
}
