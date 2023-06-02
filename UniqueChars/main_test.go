package main

import "testing"

func TestUniqueChars(t *testing.T) {
	t.Run("test if the return is true", func(t *testing.T) {
		want := true
		got := UniqueChars(TestString)
		if got != want {
			t.Errorf("Expected UniqueChars(%s) to be %t, but got %t", TestString, want, got)
		}
	})

	t.Run("test if the return is false", func(t *testing.T) {
		want := false
		got := UniqueChars(TestString)
		if got != want {
			t.Errorf("Expected UniqueChars(%s) to be %t, but got %t", TestString, want, got)
		}
	})
}

// 53.87 ns/op
func BenchmarkUniqueChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueChars(TestString)
	}
}
