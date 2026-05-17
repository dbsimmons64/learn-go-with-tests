package iterations

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeatedString := Repeat("a", 5)
	fmt.Println(repeatedString)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
