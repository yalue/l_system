A Simple L-System Library in Go
===============================

Honestly, you probably should just write your own.

Usage
-----

```
import (
	"fmt"
	"github.com/yalue/l_system"
)

func main() {
	m := l_system.NewLSystem([]byte("AB"))
	// A -> AA
	m.SetProduction('A', []byte("AA"))
	// B -> BCD
	m.SetProduction('B', []byte("BCD"))
	// C gets deleted
	m.SetProduction('C', []byte(""))
	// No production rules for 'D', it gets left as-is

	// Run two iterations.
	for i := 0; i < 2; i++ {
		err := m.Iterate()
		// Check the error
	}
	result := string(m.GetValue())
	// result should equal "AAAABCDD"
}

```

