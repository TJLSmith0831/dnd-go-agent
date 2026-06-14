package combat

/*
## How Conditions Affect Rolls
Most conditions apply to the d20 roll via Advantage or Disadvantage.
The bot's resolution path gains a step:
1. Is the attacker's condition granting disadvantage on this attack?
2. Is the defender's condition granting advantage to the attacker?
3. Do both apply? → they cancel (any advantage + any disadvantage = straight roll)
4. Roll accordingly, then check for auto-crit conditions (Paralyzed, Stunned)
*/

/*
## iota and Bit Flags
Eight conditions that can be present or absent in any combination. This
is a set of boolean flags. You could use a map[string]bool, but using
a bit flag is more efficient and easier to work with.
`iota` is a Go keyword that generates sequential numbers (0, 1, 2, 3, ...).
Each condition is assigned a unique bit position using `1 << iota`.

Used with bit-shifting, it gives each constant its own unique bit:
*/

type Condition uint64

// Each constant occupies exactly one bit. A uint16 can hold all 16 possible
// conditions in 2 bytes.
const (
	Blinded       Condition = 1 << iota // 0000_0001 = 1
	Frightened                          // 0000_0010 = 2
	Incapacitated                       // 0000_0100 = 4
	Paralyzed                           // 0000_1000 = 8
	Poisoned                            // 0001_0000 = 16
	Prone                               // 0010_0000 = 32
	Restrained                          // 0100_0000 = 64
	Stunned                             // 1000_0000 = 128
)

// Bitwise operations on the set
type ConditionSet struct {
	flags Condition
}

// Apply sets a bit: OR with the condition's bit
func (cs *ConditionSet) Apply(cond Condition) {
	cs.flags |= cond
}

// Remove clears a bit: AND-NOT (&^) with the condition's bit
func (cs *ConditionSet) Remove(cond Condition) {
	cs.flags &^= cond
}

// Has checks a bit: AND — non-zero means the bit is set
func (cs *ConditionSet) Has(cond Condition) bool {
	return cs.flags&cond != 0
}

// Does this combatant have disadvantage on their attack roll?
func (cs *ConditionSet) AttackDisadvantage() bool {
	return cs.Has(Blinded) || cs.Has(Frightened) ||
		cs.Has(Poisoned) || cs.Has(Prone) || cs.Has(Restrained)
}

// Do attacks against this combatant have advantage?
func (cs *ConditionSet) GrantsAdvantageToAttackers() bool {
	return cs.Has(Blinded) || cs.Has(Paralyzed) ||
		cs.Has(Restrained) || cs.Has(Stunned)
}
