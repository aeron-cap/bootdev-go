package arrays

import (
	"reflect"
	"testing"
)

var numbers_to_sort = []int{1, 2, 3, 4}

func TestSort(t *testing.T) {
	got := sortArray(numbers_to_sort)
	want := [][]int {{2, 4}, {1, 3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}