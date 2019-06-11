package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.ErrorF("expected '%s' but got '%s'", expected, repeated)
	}
}
