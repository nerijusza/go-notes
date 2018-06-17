package storage

// Note structure
type Note struct {
	ID          int
	Content     string
	DateCreated string
}

// Storager for storing notes
type Storager interface {
	getter
	adder
	deleter
}

type getter interface {
	// Get all notes sorted from newest to oldest
	Get() ([]Note, error)

	// Get given number newest notes sorted from newest to oldest
	GetN(int) ([]Note, error)
}

type adder interface {
	// Add new note
	Add(string) (int, error)
}

type deleter interface {
	// Delete given note
	Delete(int) error

	// Delete all notes, clear storage
	DeleteAll() error
}
