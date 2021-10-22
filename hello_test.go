package main

import "testing"

func TextHello(t *testing.T) {
	got := Hello("Neil")
	want := "Hello, Neil!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
