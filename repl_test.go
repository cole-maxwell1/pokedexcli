package main

import "testing"

func TestSanitizeInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{" ", []string{}},
		{"  ", []string{}},
		{
			"hElLo WoRlD",
			[]string{
				"hello", "world",
			}},
		{
			"TEST",
			[]string{"test"},
		},
		{
			"  tEsT  | ME",
			[]string{"test", "|", "me"},
		},
	}

	for _, c := range cases {
		got := sanitizeInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("sanitizeInput(%q) == %q, expected %q", c.input, got, c.expected)
		}
		for i := range got {
			if got[i] != c.expected[i] {
				t.Errorf("sanitizeInput(%q) == %q, expected %q", c.input, got, c.expected)
			}
		}
	}
}
