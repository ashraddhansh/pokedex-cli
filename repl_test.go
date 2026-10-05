package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HellO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello",
			expected: []string{"hello"},
		},
		{
			input: "    ",
			expected: []string{},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected){
			t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
			continue
		}

		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected, %q got %q", expectedWord, word)
			}
		}
	}

}

