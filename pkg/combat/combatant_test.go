package combat

import "testing"

func TestDeathSavingThrow(t *testing.T) {
	t.Run("error when HP above zero", func(t *testing.T) {
		pc := &PlayerCharacter{currentHP: 10}
		_, err := pc.DeathSavingThrow()
		if err == nil {
			t.Error("expected non-nil error when HP above zero, got nil")
		}
	})

	t.Run("no error when at zero HP", func(t *testing.T) {
		pc := &PlayerCharacter{currentHP: 0}
		_, err := pc.DeathSavingThrow()
		if err != nil {
			t.Errorf("expected nil error when HP at zero, got: %v", err)
		}
	})

	t.Run("eventually accumulates three failures", func(t *testing.T) {
		pc := &PlayerCharacter{currentHP: 0}
		var last DeathSaveResult
		for range 200 {
			result, err := pc.DeathSavingThrow()
			if err != nil {
				// Natural 20 woke the PC up — put them back down to keep testing.
				pc.currentHP = 0
				pc.deathSaveSuccesses = 0
				pc.deathSaveFailures = 0
				continue
			}
			last = result
			if result.Dead {
				break
			}
		}
		if !last.Dead {
			t.Errorf("expected Dead after 200 iterations, failures=%d", last.Failures)
		}
	})

	t.Run("eventually accumulates three successes", func(t *testing.T) {
		pc := &PlayerCharacter{currentHP: 0}
		var last DeathSaveResult
		for range 200 {
			result, err := pc.DeathSavingThrow()
			if err != nil {
				// Natural 20 woke the PC up — counts as stabilized, not a failure.
				// HP > 0 means the loop would error again; just stop here.
				return
			}
			last = result
			if result.Stabilized {
				break
			}
		}
		if !last.Stabilized {
			t.Errorf("expected Stabilized after 200 iterations, successes=%d", last.Successes)
		}
	})
}

func TestResolveAttack(t *testing.T) {
	t.Run("result names and AC are set correctly", func(t *testing.T) {
		attacker := &PlayerCharacter{name: "Attacker"}
		target := &Monster{name: "Target", ac: 10, currentHP: 20, maxHP: 20}

		// damageDie: 6 — needs at least 2 sides or dice.Roll errors on a hit.
		result, err := ResolveAttack(attacker, target, 0, 6, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.AttackerName != "Attacker" {
			t.Errorf("AttackerName: got %q, want %q", result.AttackerName, "Attacker")
		}
		if result.TargetName != "Target" {
			t.Errorf("TargetName: got %q, want %q", result.TargetName, "Target")
		}
		if result.TargetAC != 10 {
			t.Errorf("TargetAC: got %d, want 10", result.TargetAC)
		}
	})

	t.Run("damage is applied to target on a hit", func(t *testing.T) {
		// High attackBonus vs AC 1 — almost every roll hits.
		target := &Monster{name: "Goblin", ac: 1, currentHP: 1000, maxHP: 1000}
		attacker := &Monster{name: "Fighter"}

		hitSeen := false
		for range 100 {
			hpBefore := target.currentHP
			result, err := ResolveAttack(attacker, target, 10, 6, 0)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Hit {
				hitSeen = true
				if target.currentHP >= hpBefore {
					t.Errorf("hit but HP did not decrease: before=%d after=%d", hpBefore, target.currentHP)
				}
			}
		}
		if !hitSeen {
			t.Error("no hits in 100 attacks against AC 1 with +10 bonus — something is wrong")
		}
	})
}

func TestMiss(t *testing.T) {
	t.Run("miss deals no damage", func(t *testing.T) {
		// AC 30, attackBonus 0 — only a natural 20 hits.
		target := &Monster{name: "Ironclad", ac: 30, currentHP: 100, maxHP: 100}
		attacker := &Monster{name: "Peasant"}

		for range 100 {
			hpBefore := target.currentHP
			result, err := ResolveAttack(attacker, target, 0, 6, 0)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !result.Hit && target.currentHP != hpBefore {
				t.Errorf("miss changed target HP: before=%d after=%d", hpBefore, target.currentHP)
			}
		}
	})
}

func TestRunInitiative(t *testing.T) {
	t.Run("does not panic with empty slice", func(t *testing.T) {
		RunInitiative([]Combatant{})
	})

	t.Run("does not panic with mixed combatants", func(t *testing.T) {
		RunInitiative([]Combatant{
			&Monster{name: "Goblin"},
			&PlayerCharacter{name: "Aria"},
		})
	})

	// Test with specific combatants
	combatants := []Combatant{
		&Monster{name: "Goblin", currentHP: 7, maxHP: 7, ac: 15, dexMod: 2},
		&Monster{name: "Ogre", currentHP: 59, maxHP: 59, ac: 11, dexMod: -1},
		&PlayerCharacter{name: "Aria", currentHP: 28, maxHP: 28, ac: 16, dexMod: 3},
	}

	// Run initiative
	ordered, err := RunInitiative(combatants)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ordered) != len(combatants) {
		t.Fatalf("got %d entries, want %d", len(ordered), len(combatants))
	}
	for i := 1; i < len(ordered); i++ {
		if ordered[i].Score > ordered[i-1].Score {
			t.Errorf("not sorted at index %d: score %d > previous score %d",
				i, ordered[i].Score, ordered[i-1].Score)
		}
	}
}

// Type Assertions & Type Switches
// The combat loop has a problem: when a combatant reaches 0 HP, it needs to do different
// things depending on whether it's a monster or a player character. The Combatant interface
// doesn't provide a way to distinguish between the two types. Type assertions are Go's mechanism
// for extracting the underlying concrete value from an interface.
func TestHandleZeroHP(t *testing.T) {
	t.Run("monster dies at zero HP", func(t *testing.T) {
		m := &Monster{name: "Goblin", currentHP: 0, maxHP: 7}
		outcome := HandleZeroHP(m)
		if outcome != OutcomeDead {
			t.Errorf("got %v, want OutcomeDead", outcome)
		}
	})

	t.Run("PC at zero HP makes a death saving throw", func(t *testing.T) {
		pc := &PlayerCharacter{name: "Aria", currentHP: 0, maxHP: 28}
		outcome := HandleZeroHP(pc)
		if outcome != OutcomeDeathSave && outcome != OutcomeStabilized &&
			outcome != OutcomeDead && outcome != OutcomeAwakened {
			t.Errorf("unexpected outcome for PC: %v", outcome)
		}
	})

	t.Run("non-zero HP combatant is unaffected", func(t *testing.T) {
		m := &Monster{name: "Ogre", currentHP: 30, maxHP: 59}
		outcome := HandleZeroHP(m)
		if outcome != OutcomeAlive {
			t.Errorf("got %v, want OutcomeAlive", outcome)
		}
	})
}
