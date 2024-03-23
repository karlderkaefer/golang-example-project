package main

import (
	"testing"
)

func TestPrintVersion(t *testing.T) {
	want := "Version: development"
	got := printVersion()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
