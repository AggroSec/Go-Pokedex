package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "      HELLO WoRlD   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Pikachu CHarmander Blastoise",
			expected: []string{"pikachu", "charmander", "blastoise"},
		},
		{
			input:    "Gotta CATCHEM ALL! Pokemon!",
			expected: []string{"gotta", "catchem", "all!", "pokemon!"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths do not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match: %s %s", expectedWord, word)
			}
		}
	}
}
