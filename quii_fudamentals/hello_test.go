package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Andrei")
	want := "Hello, Andrei"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}