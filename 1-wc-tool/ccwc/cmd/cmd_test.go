package cmd

import (
	"testing"
)

const (
	testFilePath = "../../test.txt"
)

func TestCalculateBytes(t *testing.T) {
	result := calculateBytes(testFilePath)
	expected := int64(342190)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestLineCount(t *testing.T) {
	result := getLineCount(testFilePath)
	expected := 7145
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestWordCount(t *testing.T) {
	result := getWordCount(testFilePath)
	expected := 58164
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCharCount(t *testing.T) {
	// TODO: find a way to set locale during test
	result := getCharCount(testFilePath)
	expected := 339292 // output of wc on my machine
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
