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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		// add more cases here
		{
			input:    "Charmander             Bulbasaur     PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "           CharmAnder             BulbasAUR       PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actualOutput := cleanInput(c.input)
		expectedOutput := c.expected
		// Check len of the actaul slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actualOutput) != len(expectedOutput) {
			// error and continue here
			t.Errorf("Length mismatch; have: %d; want: %d", len(actualOutput), len(expectedOutput))
		}
		for i := range actualOutput {
			word := actualOutput[i]
			expectedWord := expectedOutput[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("have: %s; want: %s", word, expectedWord)
			}
		}

	}
}
