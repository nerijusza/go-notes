package config

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// Setup main config
type Setup struct {
	//Possible values: file, memory, airtable
	StorageType         string
	StorageTypeFile     FileStorageConfig
	StorageTypeAirTable AirTableConfig
}

// FileStorageConfig file storage setup
type FileStorageConfig struct {
	ProductionFile string
	TestFile       string
}

// AirTableConfig configturation for "Air Table" usage as storage engine
type AirTableConfig struct {
	Production AirTableEnvironment
	Test       AirTableEnvironment
}

// AirTableEnvironment setup for one environment
type AirTableEnvironment struct {
	Account   string
	APIKey    string
	TableName string
}

// GetSetup reads config to variable and returns it
func GetSetup() Setup {
	config.Load(file.NewSource(file.WithPath("config.yaml")))       // for main
	config.Load(file.NewSource(file.WithPath("../../config.yaml"))) // for tests
	c := Setup{}
	config.Scan(&c)
	return c
}
