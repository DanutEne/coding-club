package main

import "testing"

func TestUniqueChars(t *testing.T) {

	t.Run("uniq chars", func(t *testing.T) {
		chars := "abcde"

		got := UniqueChars(chars)
		want := true

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
