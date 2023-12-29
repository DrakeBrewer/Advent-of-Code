package parser

import (
	"testing"
)

func TestParseNums(t *testing.T) {

	got, err := ParseTxt("test.txt")
	if err != nil {
		t.Errorf("Function returned error.")
	}
	expected := 142

	if got != expected {
		t.Errorf("Expected 142; got %d", got)
	}

}
