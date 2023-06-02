package main

import "testing"

func TestIsUniqueChars(t *testing.T) {
	t.Run("test if the return is true", func(t *testing.T) {
		want := true
		got := UniqueChars(TestString)
		if got != want {
			t.Errorf("Expected isUniqueChars(%s) to be %t, but got %t", TestString, want, got)
		}
	})

	t.Run("test if the return is false", func(t *testing.T) {
		want := false
		got := UniqueChars(TestString)
		if got != want {
			t.Errorf("Expected isUniqueChars(%s) to be %t, but got %t", TestString, want, got)
		}
	})
}
