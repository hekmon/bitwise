package bitwise

import (
	"fmt"
	"math/bits"
	"strconv"
)

// Flags represents a pool of settable flags (64)
type Flags uint64

/*
	Information
*/

// Count returns the number of flags (bits) set in "pool" (Hamming weight)
func (pool Flags) Count() int {
	return bits.OnesCount64(uint64(pool))
}

// IsFlagSet returns true if the flag "flag" is present within "pool" (AND)
func (pool Flags) IsFlagSet(flag Flags) bool {
	return pool&flag == flag
}

/*
	Manipulation
*/

// SetFlag activate "flag" within "pool" (OR)
func (pool Flags) SetFlag(flag Flags) {
	pool = pool | flag
}

// ToggleFlag switch the state of "flag" within "pool" (XOR)
func (pool Flags) ToggleFlag(flag Flags) {
	pool = pool ^ flag
}

// UnsetFlag unset "flag" value within "pool" (AND NOT)
func (pool Flags) UnsetFlag(flag Flags) {
	pool = pool &^ flag
}

/*
	Printing
*/

// String implements the stdlib Stringer interface
func (pool Flags) String() string {
	return strconv.FormatUint(uint64(pool), 2)
}

// GoString implements the stdlib GoStringer interface (use "%#v")
func (pool Flags) GoString() string {
	return fmt.Sprintf("%064b(%d)", pool, pool)
}
