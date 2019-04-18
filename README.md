# Bitwise

[![GoDoc](https://godoc.org/github.com/hekmon/bitwise?status.svg)](https://godoc.org/github.com/hekmon/bitwise)

## Example

### Code

```go
package main

import (
    "fmt"
    "strconv"
    "github.com/hekmon/bitwise"
)

const (
    OptionA bitwise.Flags = 1 << iota
    OptionB
    OptionC
)

func main() {
    // Init
    fmt.Println("Initial state")
    var options bitwise.Flags
    fmt.Println("Pool:", options)
    fmt.Println("OptionA:", OptionA)
    fmt.Println("OptionB:", OptionB)
    fmt.Println("OptionC:", OptionC)
    fmt.Println()
    // Activate option B
    fmt.Println("Activate OptionB")
    options.SetFlag(OptionB)
    fmt.Println("Current pool", options)
    fmt.Println("Is OptionB on ?", options.IsFlagSet(OptionB))
    fmt.Println()
    // Toggle option A (from off to on)
    fmt.Println("Toggle OptionA")
    options.ToggleFlag(OptionA)
    fmt.Println("Current pool", options)
    fmt.Println()
    // Remove OptionB
    fmt.Println("Remove OptionB")
    options.UnsetFlag(OptionB)
    fmt.Println("Current pool", options)
    fmt.Println()
    // Something else
    fmt.Println("OptionA + OptionC debug print")
    fmt.Printf("%#v\n", OptionA+OptionC)
}
```

### Output

```raw
Initial state
Pool: 0
OptionA: 1
OptionB: 10
OptionC: 100

Activate OptionB
Current pool 10
Is OptionB on ? true

Toggle OptionA
Current pool 11

Remove OptionB
Current pool 1

OptionA + OptionC debug print
0000000000000000000000000000000000000000000000000000000000000101(5)
```
