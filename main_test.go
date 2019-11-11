package main

import "testing"

func TestGet(t *testing.T) {
	if get() == "get test" {
		t.Error("Wrong result")
	}
}
