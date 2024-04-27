package main

import "testing"

func TestName(t *testing.T) {
	got := Name("Michael")
	want := "Hello, Michael"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

