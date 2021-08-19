package l_system

import (
	"testing"
)

func TestLSystem(t *testing.T) {
	var e error
	m := NewLSystem([]byte("AC"))
	m.SetProduction('A', []byte("AA"))
	m.SetProduction('C', []byte("BC"))

	// Iterating m three times should produce 8 'A's, three additional 'B's,
	// and still end with a single 'C'.
	for i := 0; i < 3; i++ {
		t.Logf("State before iteration %d: %s\n", i+1, m.GetValue())
		e = m.Iterate()
		if e != nil {
			t.Logf("Failed iteration %d: %s\n", i+1, e)
			t.FailNow()
		}
	}
	result := string(m.GetValue())
	expected := "AAAAAAAABBBC"
	if result != expected {
		t.Logf("Expected to produce %s, but got %s\n", expected, result)
		t.FailNow()
	}

	// Test "Reset", and abide by the comment on SetProduction. Iterating this
	// should delete the 'A's and leave the 'c' as-is.
	m.Reset([]byte("AAAAAAC"))
	m.SetProduction('A', []byte(""))
	m.SetProduction('C', nil)
	e = m.Iterate()
	if e != nil {
		t.Logf("Failed iterating to remove symbols: %s\n", e)
		t.FailNow()
	}
	result = string(m.GetValue())
	expected = "C"
	if result != expected {
		t.Logf("Expected to produce %s, but got %s\n", expected, result)
		t.FailNow()
	}

	m = NewLSystem([]byte("A"))
	m.SetProduction('A', []byte("AA"))
	m.SizeLimit = 32
	for i := 0; i < 6; i++ {
		e = m.Iterate()
		if e == nil {
			continue
		}
		if i < 5 {
			t.Logf("Got unexpected error on iteration %d: %s\n", i+1, e)
			t.FailNow()
		}
		t.Logf("Got expected size-limit error on iteration %d; %s\n", i+1, e)
	}
	currentLength := uint64(len(m.GetValue()))
	if currentLength > m.SizeLimit {
		t.Logf("The length (%d bytes) of the current value exceeds the size "+
			"limit of %d bytes.\n", currentLength, m.SizeLimit)
		t.FailNow()
	}
}

// Just to verify that my example from the README works correctly.
func TestFromReadme(t *testing.T) {
	m := NewLSystem([]byte("AB"))
	m.SetProduction('A', []byte("AA"))
	m.SetProduction('B', []byte("BCD"))
	m.SetProduction('C', []byte(""))
	for i := 0; i < 2; i++ {
		e := m.Iterate()
		if e != nil {
			t.Logf("Failed iteration %d: %s\n", i+1, e)
			t.FailNow()
		}
	}
	expected := "AAAABCDD"
	result := string(m.GetValue())
	if expected != result {
		t.Logf("Expected to produce %s, but got %s\n", expected, result)
		t.FailNow()
	}
}
