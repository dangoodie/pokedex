package main

import (
	"testing"
	"strings"
)

// Test Clean Input Function
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "     hello  world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "foo    bar",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "   singleword   ",
			expected: []string{"singleword"},
		},
		{
			input:    "multiple   spaces   between   words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "   leading and trailing spaces   ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "CAPITALIZED WORDS SHOULD BE CONVERTED TO LOWERCASE",
			expected: []string{"capitalized", "words", "should", "be", "converted", "to", "lowercase"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice

		if len(actual) != len(c.expected) {
			t.Errorf("Slice length mismatched! Actual: %v Expected: %v", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.Compare(word, expectedWord) != 0 {
				t.Errorf("Words don't match! Actual: %v Expected: %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}