package array_slices_test

import (
	"fmt"
	"testing"

	array_slices "github.com/BrunoSSantana/learn-go-with-test/cmd/array_slices"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := array_slices.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := array_slices.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	sum := array_slices.Sum([]int{1, 5, 9})
	fmt.Println(sum)
	// Output: 15
}
