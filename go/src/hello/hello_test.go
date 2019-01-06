package main

import "testing"

func TestShouldReturnHelloGoMessage(t *testing.T) {
	got := Hello()
	want := "Hello Go"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
