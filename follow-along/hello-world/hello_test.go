package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Josh")
	want := "Hello, Josh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
