package storage

import (
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	// inittialize memory storage
	s := Memory{}
	s.Init()

	// pass to dedicated memory tester, to test against storager interface
	tester := Tester{&s, t}
	tester.Test()
}
