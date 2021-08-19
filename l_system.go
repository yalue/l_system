// The l_system package defines a library for generating strings using an
// L-system.  See https://en.wikipedia.org/wiki/L-system.
package l_system

import (
	"fmt"
)

// The default limit on the size of the produced string: 100 MB.
const DefaultSizeLimit = 1024 * 1024 * 100

// Returned by Iterate() if the size limit of the produced string would
// increase past the specified SizeLimit.
var SizeLimitExceededError = fmt.Errorf("Further iteration would exceed " +
	"the system's current limit on the size of the produced string.")

// Holds the production rules in the L-system as well as its current state.
type LSystem struct {
	// A slice holding each production, indexed by the byte they match and
	// containing the byte slice the matched byte should be replaced with.
	productions [][]byte
	// The current string of the L-system
	currentState []byte
	// The limit on the number of bytes that can be produced by the L-system.
	SizeLimit uint64
}

// Returns a new slice holding a copy of b.
func copyBytes(b []byte) []byte {
	toReturn := make([]byte, len(b))
	copy(toReturn, b)
	return toReturn
}

// Sets the production rule for the given byte, overwriting any previously-set
// production for the byte. Copies the "replacement" slice. If 'replacement' is
// nil, the production is removed entirely, meaning that the character is left
// as-is. If it is instead an empty slice, then the production will remove the
// matched symbol.
func (m *LSystem) SetProduction(matches byte, replacement []byte) {
	if replacement == nil {
		m.productions[matches] = nil
		return
	}
	m.productions[matches] = copyBytes(replacement)
}

// Resets the L-systems current state to the given bytes, without clearing the
// production rules or changing the size limit. Copies the new initial value,
// so the caller won't have access to the L-system's internal buffer.
func (m *LSystem) Reset(initialValue []byte) {
	m.currentState = copyBytes(initialValue)
}

// Returns the size of the resulting string, after applying a single iteration
// of productions.
func (m *LSystem) getNextSize() uint64 {
	toReturn := uint64(0)
	var p []byte
	for _, c := range m.currentState {
		p = m.productions[c]
		if p == nil {
			// A nil production means that the symbol remains unchanged.
			toReturn++
			continue
		}
		toReturn += uint64(len(p))
	}
	return toReturn
}

// Applies the system's production rules exactly once, updating the current
// string. If an error occurs, this returns an error and doesn't change the
// current string at all.
func (m *LSystem) Iterate() error {
	newSize := m.getNextSize()
	if newSize > m.SizeLimit {
		return SizeLimitExceededError
	}
	newState := make([]byte, 0, newSize)
	var p []byte
	for _, c := range m.currentState {
		p = m.productions[c]
		if p == nil {
			newState = append(newState, c)
			continue
		}
		newState = append(newState, p...)
	}
	m.currentState = newState
	return nil
}

// Returns the current string generated by the L-system.
func (m *LSystem) GetValue() []byte {
	return m.currentState
}

// Creates a new L-system instance, with no production rules, and the state
// set to the given initial value. Copies the initial value, so the caller
// won't have access to the L-system's internal buffer.
func NewLSystem(initialValue []byte) *LSystem {
	return &LSystem{
		productions:  make([][]byte, 256),
		currentState: copyBytes(initialValue),
		SizeLimit:    DefaultSizeLimit,
	}
}