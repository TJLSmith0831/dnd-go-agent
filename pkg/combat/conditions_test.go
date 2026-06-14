package combat

import "testing"

func TestConditionSet(t *testing.T) {
	t.Run("Apply and Has round-trip", func(t *testing.T) {
		var cs ConditionSet
		cs.Apply(Poisoned)
		if !cs.Has(Poisoned) {
			t.Error("expected Poisoned after Apply")
		}
	})

	t.Run("Remove clears the condition", func(t *testing.T) {
		var cs ConditionSet
		cs.Apply(Prone)
		cs.Remove(Prone)
		if cs.Has(Prone) {
			t.Error("expected Prone cleared after Remove")
		}
	})

	t.Run("Remove does not affect other conditions", func(t *testing.T) {
		var cs ConditionSet
		cs.Apply(Poisoned)
		cs.Apply(Prone)
		cs.Remove(Prone)
		if !cs.Has(Poisoned) {
			t.Error("Remove(Prone) should not clear Poisoned")
		}
	})

	t.Run("Poisoned gives attack disadvantage", func(t *testing.T) {
		var cs ConditionSet
		cs.Apply(Poisoned)
		if !cs.AttackDisadvantage() {
			t.Error("Poisoned should give attack disadvantage")
		}
	})

	t.Run("Paralyzed grants advantage to attackers", func(t *testing.T) {
		var cs ConditionSet
		cs.Apply(Paralyzed)
		if !cs.GrantsAdvantageToAttackers() {
			t.Error("Paralyzed should grant advantage to attackers")
		}
	})
}
