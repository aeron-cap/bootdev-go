package arrays

import (
	"reflect"
	"testing"
)

var numbers = []int{1, 2, 3}
var numbers_1 = []int{1, 2, 3}

func TestSum(t *testing.T) {
	got := ArraySum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		ArraySum(numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll(numbers, numbers_1)
	want := []int{6, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll(numbers, numbers_1)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails(numbers, numbers_1)
	want := []int{5, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}