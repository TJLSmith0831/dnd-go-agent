package referee

import (
	"errors"
	"fmt"
	"strings"
)

type Skill string

const (
	SkillStealth    Skill = "stealth"
	SkillPerception Skill = "perception"
	SkillPersuasion Skill = "persuasion"
	SkillAthletics  Skill = "athletics"
	SkillArcana     Skill = "arcana"
)

var ErrUnknownSkill = errors.New("unknown skill")

type RefereeError struct {
	Input  string
	Reason string
}

func (e *RefereeError) Error() string {
	return fmt.Sprintf("referee error for input %q: %v", e.Input, e.Reason)
}

func (e *RefereeError) Unwrap() error {
	return ErrUnknownSkill
}

// keywords maps input fragments to skills
var keywords = map[string]Skill{
	"sneak":    SkillStealth,
	"hide":     SkillStealth,
	"notice":   SkillPerception,
	"spot":     SkillPerception,
	"convince": SkillPersuasion,
	"persuade": SkillPersuasion,
	"climb":    SkillAthletics,
	"jump":     SkillAthletics,
	"recall":   SkillArcana,
	"magic":    SkillArcana,
}

func ResolveSkill(input string) (Skill, error) {
	lower := strings.ToLower(input)
	for kw, skill := range keywords {
		if strings.Contains(lower, kw) {
			return skill, nil
		}
	}
	return "", &RefereeError{Input: input, Reason: "no keyword matched"}
}
