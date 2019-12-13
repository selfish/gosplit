package gosplit

import (
	"testing"
)

func TestAbTest(t *testing.T) {
	expected := "control"
	if actual := abTest(); actual != expected {
		t.Errorf("Hello() = %q, want %q", actual, expected)
	}
}
