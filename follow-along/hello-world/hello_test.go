package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Josh")
		want := "Hello, Josh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world when an empty string is provided", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Telling go to treat this function as a helper function, so that when it fails, it will be
	// reported as a failure on the line number in our test where it was called, rather than inside our helper function.

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
