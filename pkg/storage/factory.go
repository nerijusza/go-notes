package storage

// Factory returns initialized storage for use by config (will be implemented later, now just memory)
type Factory struct {
	storage Storager
}

// Get returns storage implementation
func (t *Factory) Get() (*Storager, error) {
	err := t.initialize()
	if err != nil {
		return nil, err
	}
	return &t.storage, nil
}

func (t *Factory) initialize() error {
	if t.storage != nil {
		return nil
	}

	return t.initializeMemoryStorage()
}

func (t *Factory) initializeMemoryStorage() error {
	s := Memory{}
	s.Init()
	t.storage = &s

	return nil
}

func (t *Factory) initializeFileStorage() error {
	s := FileStorage{}
	err := s.Init("website_data.txt")
	if err == nil {
		t.storage = &s
	}
	return err
}
