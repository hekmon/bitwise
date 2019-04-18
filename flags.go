package bitwise

import (
	"fmt"
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
	return fmt.Sprintf("dec(%d)|bin(%064b)", f, f)
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
