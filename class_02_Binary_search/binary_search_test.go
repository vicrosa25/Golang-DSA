package class_02_Binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("The key is in the slice", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Search(nums, 9)
		want := 9
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("The key is not in the slice", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Search(nums, 10)
		want := -1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
