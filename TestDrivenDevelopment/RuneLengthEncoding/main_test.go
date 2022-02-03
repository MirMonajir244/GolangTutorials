package main

import "testing"

func TestEncode(t *testing.T) {
	for _, c := range test_cases {
		if c.expected == Encode(c.input) {
			t.Logf("Pass for %s", c.description)
		} else {
			t.Fatalf("Fail : %s , wanted : %s, got: %s", c.description, c.expected, Encode(c.input))
		}
	}
}

func TestDigit(t *testing.T) {
	if isDigit("5") {
		t.Logf("Pass digit test %d", 5)
	} else {
		t.Fatalf("Fail : digit test , wanted : true, got: false")
	}
}
