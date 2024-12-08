package main

import "testing"

func TestCleanInput( t *testing.T) {
	cases := []struct{
		input string
		expected []string
	} {
		{
			input: "hello world",
			expected: []string {
				"hello",
				"world",
			},
		},
	}

	for _, case := range cases {
		actual := CleanInput(case.input)
		if len(actual) != len(case.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(case.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := case.expected[i]
		}
	}
}