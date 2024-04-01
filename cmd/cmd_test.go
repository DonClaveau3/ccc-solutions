package cmd

import (
	"fmt"
	"os"
	"testing"
)

const (
	testFilePath = "../test.txt"
)

func TestScanForStats(t *testing.T) {
	f, err := os.Open(testFilePath)
	if err != nil {
		fmt.Println(err)
	}
	b, c, w, l, err := scanForStats(f)
	if err != nil {
		fmt.Println(err)
	}
	expected_bytes := 342190
	if b != expected_bytes {
		t.Errorf("Bytes: expected %d, but got %d", expected_bytes, b)
	}
	expected_chars := 339292
	if c != expected_chars {
		t.Errorf("Characters: expected %d, but got %d", expected_chars, c)
	}
	expected_words := 58164
	if w != expected_words {
		t.Errorf("Words: expected %d, but got %d", expected_words, w)
	}
	expected_lines := 7145
	if l != expected_lines {
		t.Errorf("Lines: expected %d, but got %d", expected_lines, l)
	}
}
