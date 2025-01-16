package class_01_Linear_search

import "testing"

func TestLinearSearch(t *testing.T) {
	numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	got := Search(numbers, 23)
	want := true

	if got != want {
		t.Errorf("LinearSearch got %v, want %v", got, want)
	}
}
