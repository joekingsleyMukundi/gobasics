package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but got none")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Error("expected", tt.expected, "but got", got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10, 1.0)
	if err != nil {
		t.Error("Got an error whaen we should not a have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0.0)
	if err == nil {
		t.Error("Did not get an error whaen we should have")
	}
}
