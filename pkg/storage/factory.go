package storage

import (
	"errors"

	"github.com/nerijusza/go-notes/pkg/config"
)

// Factory returns initialized storage for use by config (will be implemented later, now just memory)
type Factory struct {
	storage Storager
}

// Get returns storage implementation
func (t *Factory) Get() (Storager, error) {
	err := t.initialize()
	if err != nil {
		return nil, err
	}
	return t.storage, nil
}

func (t *Factory) initialize() error {
	if t.storage != nil {
		return nil
	}

	setup := config.GetSetup()

	if setup.StorageType == "file" {
		return t.initializeFileStorage(setup.StorageTypeFile)
	}
	if setup.StorageType == "memory" {
		return t.initializeMemoryStorage()
	}

	return errors.New("Unknown memory storage type: " + setup.StorageType)
}

func (t *Factory) initializeMemoryStorage() error {
	s := Memory{}
	s.Init()
	t.storage = &s

	return nil
}

func (t *Factory) initializeFileStorage(setup config.FileStorageConfig) error {
	s := FileStorage{setup.Directory + "/" + setup.File}
	err := s.Init(setup.File)
	if err == nil {
		t.storage = &s
	}
	return err
}
