package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	expected := "Hello, world!"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
