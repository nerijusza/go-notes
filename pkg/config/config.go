package config

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// Setup main config
type Setup struct {
	//Possible values: file, memory
	StorageType     string
	StorageTypeFile FileStorageConfig
}

// FileStorageConfig file storage setup
type FileStorageConfig struct {
	Directory string
	File      string
}

// GetSetup reads config to variable and returns it
func GetSetup() Setup {
	config.Load(file.NewSource(file.WithPath("config.yaml")))
	c := Setup{}
	config.Scan(&c)
	return c
}
