package main

import "testing"

	got := Hello("Neil")
	want := "Hello, Neil!"
func TestHello(t *testing.T) {

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
