package arraysandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of slices numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{8, 9, 10})
		want := []int{3, 27}

		if !slices.Equal(got, want) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

func TestSumTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
	t.Run("sum tails for a collection of slices numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{8, 9, 10})
		want := []int{5, 19}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
