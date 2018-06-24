package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/nerijusza/go-notes/pkg/helper"
)

// FileStorage implementation of StorageInterface in file system
type FileStorage struct {
	fileName string // file name to work with, just name not absolute or partial paths
}

// Init initializes storage aka constructor
func (t *FileStorage) Init() error {
	return t.createFile()
}

func (t FileStorage) createFile() error {
	// detect if file exists
	var _, err = os.Stat(t.fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		_, err = os.Create(t.fileName)
	}
	return err
}

// Get gets all notes sorted fron newest to oldest
func (t *FileStorage) Get() ([]Note, error) {
	notes, err := t.getNotesFromFile()
	if err != nil {
		return nil, err
	}

	return reverseNotesArray(notes), nil
}

// GetN get newest n notes
func (t *FileStorage) GetN(quantity int) ([]Note, error) {
	notes, err := t.Get()
	if err != nil {
		return nil, err
	}

	if len(notes) > quantity {
		notes = notes[0:quantity]
	}

	return notes, nil
}

// Add saves given string as note
func (t *FileStorage) Add(content string) (int, error) {
	notes, err := t.getNotesFromFile()
	if err != nil {
		return 0, err
	}

	biggestID, err := t.getBiggestNoteID()
	if err != nil {
		return 0, err
	}

	notes = append(notes, Note{biggestID + 1, content, helper.GetCurrentTime()})

	err = t.saveNotesToFile(notes)
	if err != nil {
		return 0, err
	}

	return biggestID + 1, nil
}

// Delete deletes given note by id
func (t *FileStorage) Delete(ID int) error {
	notes, err := t.getNotesFromFile()
	if err != nil {
		return err
	}

	for i, note := range notes {
		if note.ID == ID {
			notes = append(notes[:i], notes[i+1:]...)
			return t.saveNotesToFile(notes)
		}
	}
	return nil
}

// DeleteAll deletes all notes from storage
func (t *FileStorage) DeleteAll() error {
	return t.saveNotesToFile(make([]Note, 0))
}

func (t *FileStorage) getNotesFromFile() ([]Note, error) {
	fileContent, err := ioutil.ReadFile(t.fileName)
	if err != nil {
		return nil, err
	}

	notes := make([]Note, 0)
	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &notes)
	}

	return notes, err
}

func (t *FileStorage) saveNotesToFile(notes []Note) error {
	fileContent, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(t.fileName, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (t *FileStorage) getBiggestNoteID() (int, error) {
	notes, err := t.getNotesFromFile()
	if err != nil {
		return 0, err
	}

	if len(notes) == 0 {
		return 0, nil
	}

	return notes[len(notes)-1].ID, nil
}
