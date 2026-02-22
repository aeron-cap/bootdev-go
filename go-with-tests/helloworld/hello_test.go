package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to a person", func(t *testing.T) {
		got := Hello("Aeron", "English")
		want := "Hello Aeron"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello when theres no provided name", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}