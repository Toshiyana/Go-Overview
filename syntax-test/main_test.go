package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	// test-data
	{"valid-data", 100, 2, 50, false},
	{"invalid-data", 100, 0, 0, true},
	{"expect-5", 100, 20, 5, true},
}

// execute test-command in terminal
// * "go test" or "go test -v": check Fail or Path
// * "go test -cover": can check coverage in tests
// * "go test -coverprofile=coverage.out && go tool cover -html=coverage.out"
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f, but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
