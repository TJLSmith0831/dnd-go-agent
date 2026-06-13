package dice

import (
	"fmt"
	"math"
	"math/rand"
)

// Roll returns a random number from 1 to sides, inclusive.
func Roll(sides int) (int, error) {
	if sides < 2 {
		return 0, fmt.Errorf("Roll: need at least 2 sides, got %d", sides)
	}
	return rand.Intn(sides) + 1, nil
}

// AbilityModifier converts a D&D 5e ability score to its modifier.
// Uses math.Floor because Go's integer division truncates toward zero,
// which gives the wrong result for odd scores below 10 (e.g. score 9 → 0, not −1).
func AbilityModifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// CheckResult holds the outcome of a D&D 5e ability check.
type CheckResult struct {
	Roll     int
	Modifier int
	Total    int
	DC       int
	Success  bool
}

// AbilityCheck resolves a D&D 5e ability check against a Difficulty Class.
func AbilityCheck(score, dc int) (CheckResult, error) {
	roll, err := Roll(20)
	if err != nil {
		return CheckResult{}, err
	}
	mod := AbilityModifier(score)
	total := roll + mod
	return CheckResult{
		Roll:     roll,
		Modifier: mod,
		Total:    total,
		DC:       dc,
		Success:  total >= dc,
	}, nil
}
