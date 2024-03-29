package main

import (
	"fmt"
	"testing"
)

func TestPrintVersion(t *testing.T) {
	want := fmt.Sprintf("my app %s, commit %s, built at %s", version, commit, date)
	got := printVersion()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
