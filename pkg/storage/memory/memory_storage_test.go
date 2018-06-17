package storage

import (
	"testing"

	"github.com/nerijusza/go-notes/pkg/storage"
)

func TestMemoryStorage(t *testing.T) {
	// inittialize memory storage
	s := Memory{}
	s.Init()

	// pass to dedicated memory tester, to test against storager interface
	tester := storage.Tester{&s, t}
	tester.Test()
}
