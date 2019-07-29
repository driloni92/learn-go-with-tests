package main

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, world"
	got := greetWorld()

	if want != got {
		t.Errorf("Expected %q, got %q\n", want, got)
	}
}

func TestHelloBob(t *testing.T) {
	want := "Hello, Bob"
	got := greet(name + " ,")

	if want != got {
		t.Errorf("Expected %q, got %q\n", want, got)
	}

}
