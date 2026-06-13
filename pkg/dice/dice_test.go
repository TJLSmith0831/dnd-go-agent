package dice

import "testing"

// Go's test runner looks for files named `*_test.go` and functions named `Test*`
// The files are compiled and run with the `go test` command.

// CONVENTION: The test file for dice.go lives in the same directory and
// is named dice_test.go
// func TestAbilityModifier(t *testing.T) {
// 	got := AbilityModifier(10)
// 	if got != 0 {
// 		t.Errorf("score 10: got %d, want 0", got)
// 	}
// }

// The table-driven pattern
// One test function per case doesn't scale. Go's idiomatic answer is to define
// a slice of test cases and iterate over them. This is so common it's practically
// a law.
func TestAbilityModifier(t *testing.T) {

	// The struct type is defined inline. This is common practice in Go tests.
	tests := []struct {
		name  string
		score int
		want  int
	}{
		// Named subtests with t.Run
		{"min score", 1, -5},
		{"low score", 8, -1},
		{"floor-division case", 9, -1}, // the floor-division case — would be 0 with plain / in Go
		{"mid score", 10, 0},
		{"upper mid", 11, 0},
		{"high score", 14, 2},
		{"max score", 20, 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AbilityModifier(tc.score)
			if got != tc.want {
				t.Errorf("score %d: got %d, want %d", tc.score, got, tc.want)
			}
		})
	}
}

// Random output is harder to test than deterministic functions.
// You can't assert the exact value. Instead, test the contract:
// - The result is within the expected range
// - The function behaves consistently (e.g., same inputs produce same outputs)
// A common approach is to test the range of possible outputs.
func TestRoll(t *testing.T) {
	t.Run("valid range", func(t *testing.T) {
		for range 1000 {
			got, err := Roll(20)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got < 1 || got > 20 {
				t.Errorf("got %d, want value between 1 and 20", got)
			}
		}
	})

	t.Run("invalid sides", func(t *testing.T) {
		_, err := Roll(0)
		if err == nil {
			t.Error("expected error for sides=0, got nil")
		}
	})
}
