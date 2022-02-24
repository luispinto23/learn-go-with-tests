package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Luís")
	expected := "Hello, Luís!"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
