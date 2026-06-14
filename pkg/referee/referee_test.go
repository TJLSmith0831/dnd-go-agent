package referee

import (
	"errors"
	"testing"
)

func TestResolveSkill(t *testing.T) {
	t.Run("recognizes stealth", func(t *testing.T) {
		skill, err := ResolveSkill("I try to sneak past the guard")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if skill != SkillStealth {
			t.Errorf("got %v, want Stealth", skill)
		}
	})

	t.Run("returns ErrUnknownSkill for unrecognised input", func(t *testing.T) {
		_, err := ResolveSkill("I want to do a cool thing")
		if !errors.Is(err, ErrUnknownSkill) {
			t.Errorf("expected ErrUnknownSkill, got %v", err)
		}
	})

	t.Run("RefereeError carries the original input", func(t *testing.T) {
		_, err := ResolveSkill("I want to do a cool thing")
		var refErr *RefereeError
		if !errors.As(err, &refErr) {
			t.Fatalf("expected *RefereeError in chain, got %T", err)
		}
		if refErr.Input != "I want to do a cool thing" {
			t.Errorf("Input: got %q", refErr.Input)
		}
	})
}
