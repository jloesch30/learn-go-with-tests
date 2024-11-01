package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Josh", "")
		want := "Hello, Josh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Josh", "French")
		want := "Bonjour, Josh"
		assertCorrectMessage(t, got, want)
	})
}

// we can shorten the syntax of the test function by ommiting the "string" type in the function signature, since they are the same
func assertCorrectMessage(t testing.TB, got, want string) {
	// Telling go to treat this function as a helper function, so that when it fails, it will be
	// reported as a failure on the line number in our test where it was called, rather than inside our helper function.

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
