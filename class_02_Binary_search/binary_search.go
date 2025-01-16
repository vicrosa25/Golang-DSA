package class_02_Binary_search

// Search Low is included High is excluded
func Search(nums []int, key int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == key {
			return mid
		} else if nums[mid] > key {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
