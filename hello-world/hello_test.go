package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Luís", "Spanish")
		want := "Hola, Luís"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in english", func(t *testing.T) {
		got := Hello("Luís", "")
		want := "Hello, Luís"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}
