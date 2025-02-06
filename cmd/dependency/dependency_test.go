package dependency

import (
	"bytes"
	"testing"
)

func TestGreetService(t *testing.T) {
	buffer := bytes.Buffer{}
	greet := &Greet{}
	service := NewGreetService(greet)

	service.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
