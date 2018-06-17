package storage

import (
	"github.com/nerijusza/go-notes/pkg/helper"
	"github.com/nerijusza/go-notes/pkg/storage"
)

// Memory type implementation od StorageInterface
type Memory struct {
	// map for storing notes
	table        map[int]storage.Note
	biggestIndex int
}

// Init initializes storage aka constructor
func (t *Memory) Init() {
	t.table = make(map[int]storage.Note)
}

// Get gets all notes sorted fron newest to oldest
func (t Memory) Get() ([]storage.Note, error) {
	var list []storage.Note

	// reversed order from newest to oldest
	for i := range t.table {
		list = append(list, t.table[i])
	}
	return storage.SortByIDDesc(list), nil
}

// GetN get newest n notes
func (t Memory) GetN(quantity int) ([]storage.Note, error) {
	list, err := t.Get()
	if len(list) > quantity {
		list = list[0:quantity]
	}
	return list, err
}

// Add saves given string as note
func (t *Memory) Add(content string) (int, error) {
	t.biggestIndex++
	note := storage.Note{t.biggestIndex, content, helper.GetCurrentTime()}
	t.table[t.biggestIndex] = note
	return t.biggestIndex, nil
}

// Delete deletes given note by id
func (t *Memory) Delete(id int) error {
	delete(t.table, id)
	return nil
}

// DeleteAll deletes all notes from storage
func (t *Memory) DeleteAll() error {
	t.table = make(map[int]storage.Note)
	return nil
}
