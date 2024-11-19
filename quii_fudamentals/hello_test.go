package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("say hello to a person", func(t *testing.T) {
    got := Hello("Andrei")
    want := "Hello, Andrei"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })	
	got := Hello("Andrei")
	want := "Hello, Andrei"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
