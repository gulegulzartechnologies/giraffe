package giraffe

import (
	"testing"
)

func TestHasher(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hash of 'hello'", "hello", "5d41402abc4b2a76b9719d911017c592"},
		{"Hash of 'world'", "world", "7d793037a0760186574b0282f2f435e7"},
		{"Hash of empty string", "", "d41d8cd98f00b204e9800998ecf8427e"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hasher(test.input)
			if result != test.expected {
				t.Errorf("Hasher(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
