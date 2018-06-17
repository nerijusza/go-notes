package storage

// Note structure
type Note struct {
	ID          int
	Content     string
	DateCreated string
}

// Storager for storing notes
type Storager interface {
	Getter
	Adder
	Deleter
}

// Getter to get notes from storage
type Getter interface {
	// Get all notes sorted from newest to oldest
	Get() ([]Note, error)

	// Get given number newest notes sorted from newest to oldest
	GetN(int) ([]Note, error)
}

// Adder for adding notes to storage
type Adder interface {
	// Add new note
	Add(string) (int, error)
}

// Deleter to delete notes from storage
type Deleter interface {
	// Delete given note
	Delete(int) error

	// Delete all notes, clear storage
	DeleteAll() error
}
