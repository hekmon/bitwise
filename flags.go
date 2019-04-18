package bitwise

import (
	"fmt"
	"math/bits"
	"strconv"
)

// Flags represents a pool of settable flags (64)
type Flags uint64

// String implements the stdlib Stringer interface
func (f Flags) String() string {
	return strconv.FormatUint(uint64(f), 2)
}

// GoString implements the stdlib GoStringer interface (use "%#v")
func (f Flags) GoString() string {
	return fmt.Sprintf("%064b(%d)", f, f)
}

// IsFlagSet returns true if the flag "flag" is present within "pool" (AND)
func IsFlagSet(pool, flag Flags) bool {
	return pool&flag == flag
}

// UnsetFlag unset "flag" value within "pool" (AND NOT)
func UnsetFlag(pool, flag Flags) Flags {
	return pool &^ flag
}

// SetFlag activate "flag" within "pool" (OR)
func SetFlag(pool, flag Flags) Flags {
	return pool | flag
}

// ToggleFlag switch the state of "flag" within "pool" (XOR)
func ToggleFlag(pool, flag Flags) Flags {
	return pool ^ flag
}

// CountFlags returns the number of flags (bits) set in "pool" (Hamming weight)
func CountFlags(pool Flags) int {
	return bits.OnesCount64(uint64(pool))
}
