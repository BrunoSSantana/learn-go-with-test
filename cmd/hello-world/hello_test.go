package helloworld_test

import (
	"testing"

	helloworld "github.com/BrunoSSantana/learn-go-with-test/cmd/hello-world"
)

func TestHello(t *testing.T) {
	got := helloworld.Hello("Bruno")
	want := "Hello, Bruno"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
