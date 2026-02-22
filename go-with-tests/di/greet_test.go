package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aeron")

	got := buffer.String()
	want := "Hello, Aeron"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}