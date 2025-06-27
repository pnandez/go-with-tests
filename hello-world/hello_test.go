package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Pablo")
	want := "Helo, Pablo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
