package combat

import (
	"fmt"
	"sort"

	"github.com/tjlsmith0831/dnd-go-agent/pkg/dice"
)

// -------------------------
// ------- Interface -------
// -------------------------

// Combatant is anything that can participate in D&D 5e combat.
type Combatant interface {
	Name() string
	CurrentHP() int
	MaxHP() int
	AC() int
	InitiativeBonus() int // DEX modifier
	IsAlive() bool
	TakeDamage(amount int)
}

// DeathSaveResult holds the outcome of a single death saving throw.
type DeathSaveResult struct {
	Roll       int
	Successes  int
	Failures   int
	Stabilized bool // three successes reached
	Dead       bool // three failures reached
	Awakened   bool // natural 20: regain 1 HP
}

// Monster represents a D&D 5e monster in combat.
type Monster struct {
	name      string
	currentHP int
	maxHP     int
	ac        int
	dexMod    int
}

func (m *Monster) Name() string         { return m.name }
func (m *Monster) CurrentHP() int       { return m.currentHP }
func (m *Monster) MaxHP() int           { return m.maxHP }
func (m *Monster) AC() int              { return m.ac }
func (m *Monster) InitiativeBonus() int { return m.dexMod }
func (m *Monster) IsAlive() bool        { return m.currentHP > 0 }
func (m *Monster) TakeDamage(amount int) {
	m.currentHP -= amount
	if m.currentHP < 0 {
		m.currentHP = 0
	}
}

// PlayerCharacter represents a D&D 5e player character in combat.
type PlayerCharacter struct {
	name               string
	currentHP          int
	maxHP              int
	ac                 int
	dexMod             int
	deathSaveSuccesses int
	deathSaveFailures  int
	dead               bool
}

func (pc *PlayerCharacter) Name() string         { return pc.name }
func (pc *PlayerCharacter) CurrentHP() int       { return pc.currentHP }
func (pc *PlayerCharacter) MaxHP() int           { return pc.maxHP }
func (pc *PlayerCharacter) AC() int              { return pc.ac }
func (pc *PlayerCharacter) InitiativeBonus() int { return pc.dexMod }
func (pc *PlayerCharacter) IsAlive() bool        { return pc.currentHP > 0 }
func (pc *PlayerCharacter) TakeDamage(amount int) {
	pc.currentHP -= amount
	if pc.currentHP < 0 {
		pc.currentHP = 0
	}
}

// -------------------------
// ------- Actions ---------
// -------------------------

// DeathSavingThrow resolves a D&D 5e death saving throw for a downed PC.
// Returns an error if the PC is not at 0 HP.
func (pc *PlayerCharacter) DeathSavingThrow() (DeathSaveResult, error) {
	if pc.currentHP > 0 {
		return DeathSaveResult{}, fmt.Errorf("DeathSavingThrow: %s is not at 0 HP", pc.name)
	}

	roll, err := dice.Roll(20)
	if err != nil {
		return DeathSaveResult{}, err
	}

	switch {
	case roll == 20:
		// Natural 20: wake up with 1 HP, reset counters
		pc.currentHP = 1
		pc.deathSaveSuccesses = 0
		pc.deathSaveFailures = 0
		pc.dead = false
	case roll == 1:
		// Natural 1: counts as two failures
		pc.deathSaveFailures += 2
	case roll >= 10:
		pc.deathSaveSuccesses++
	default:
		pc.deathSaveFailures++
	}

	return DeathSaveResult{
		Roll:       roll,
		Successes:  pc.deathSaveSuccesses,
		Failures:   pc.deathSaveFailures,
		Stabilized: pc.deathSaveSuccesses >= 3,
		Dead:       pc.deathSaveFailures >= 3,
		Awakened:   roll == 20,
	}, nil
}

// -------------------------
// ------- Engine  ---------
// -------------------------

// The engine works on the interface. Monster and PlayerCharacter
// can be mixed freely in the same slice.
type InitiativeEntry struct {
	Combatant Combatant
	Score     int
}

func RunInitiative(combatants []Combatant) ([]InitiativeEntry, error) {
	entries := make([]InitiativeEntry, len(combatants))

	// Roll initiative for each combatant
	for i, c := range combatants {
		roll, err := dice.Roll(20)
		if err != nil {
			// Handle error appropriately
			return nil, err
		}
		entries[i] = InitiativeEntry{
			Combatant: c,
			Score:     roll + c.InitiativeBonus(),
		}
	}

	// Sort by initiative (highest first)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score > entries[j].Score
	})

	return entries, nil
}

// AttackResult holds the full outcome of a single attack.
type AttackResult struct {
	AttackerName string
	TargetName   string
	Roll         int
	AttackTotal  int
	TargetAC     int
	Hit          bool
	Critical     bool
	Damage       int // 0 on a miss
}

// ResolveAttack resolves a D&D 5e attack roll and applies damage on a hit.
// attackBonus = ability modifier + proficiency bonus for this weapon.
// damageDie   = number of sides on the weapon's damage die (e.g. 8 for 1d8).
// damageBonus = ability modifier added to damage (same one used for the attack roll).
func ResolveAttack(attacker, target Combatant, attackBonus, damageDie, damageBonus int) (AttackResult, error) {
	roll, err := dice.Roll(20)
	if err != nil {
		return AttackResult{}, err
	}

	critical := roll == 20
	// Natural 20 always hits; natural 1 always misses.
	hit := critical || (roll != 1 && roll+attackBonus >= target.AC())

	result := AttackResult{
		AttackerName: attacker.Name(),
		TargetName:   target.Name(),
		Roll:         roll,
		AttackTotal:  roll + attackBonus,
		TargetAC:     target.AC(),
		Hit:          hit,
		Critical:     critical,
	}

	if hit {
		dmg, err := dice.Roll(damageDie)
		if err != nil {
			return AttackResult{}, err
		}
		if critical {
			// Critical hit: roll the damage die a second time and add it.
			extra, err := dice.Roll(damageDie)
			if err != nil {
				return AttackResult{}, err
			}
			dmg += extra
		}
		result.Damage = dmg + damageBonus
		target.TakeDamage(result.Damage)
	}

	return result, nil
}

// Type Assertions have two forms
// Single-value form — panics if c is not  *PlayerCharacter
// 		pc := c.(*PlayerCharacter)
// Two-value form — safe, ok is false if the assertion fails
//		pc, ok := c.(*PlayerCharacter)
//		if ok {
//		    // pc is a *PlayerCharacter here
//		}
// Always use the two-value form when you're not sure the type is correct
// The single-valuye is rare in production code and mostly just for testing

/**
Type Switches
When you need to branch on more than one possible type, a type switch is cleaner than chaining assertions
switch v := c.(type) {
case *PlayerCharacter:
	// v is *PlayerCharacter here
	v.DeathSavingThrow()
case *Monster:
	// v is *Monster here
	fmt.Printf("%s is slain\n", v.Name())
default:
	// v is Combatant — the original interface type
}
}
*/

// type ZeroHPOutcome string creates a new type — you cannot accidentally assign a plain string to a
// ZeroHPOutcome variable without an explicit conversion. This is intentional: it prevents passing "ded"
// where OutcomeDead is expected. It's a lightweight alternative to an enum.
type ZeroHPOutcome string

const (
	OutcomeAlive      ZeroHPOutcome = "alive"
	OutcomeDead       ZeroHPOutcome = "dead"
	OutcomeDeathSave  ZeroHPOutcome = "death_save"
	OutcomeStabilized ZeroHPOutcome = "stabilized"
	OutcomeAwakened   ZeroHPOutcome = "awakened"
)

func HandleZeroHP(c Combatant) ZeroHPOutcome {
	if c.CurrentHP() > 0 {
		return OutcomeAlive
	}

	// Create type switch dependent on the combatant type
	switch v := c.(type) {
	case *PlayerCharacter:
		result, err := v.DeathSavingThrow()
		if err != nil {
			return OutcomeDead // shouldn't happen — HP check above guards it
		}
		switch {
		case result.Awakened:
			return OutcomeAwakened
		case result.Dead:
			return OutcomeDead
		case result.Stabilized:
			return OutcomeStabilized
		default:
			return OutcomeDeathSave
		}
	default:
		// Monsters and anything else: dead at 0 HP
		return OutcomeDead
	}
}
