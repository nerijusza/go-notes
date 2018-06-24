package storage

import (
	"testing"

	"github.com/nerijusza/go-notes/pkg/config"
)

func TestAirTableStorage(t *testing.T) {
	setup := config.GetSetup()
	s := AirTableStorage{setup.StorageTypeAirTable.Test.Account, setup.StorageTypeAirTable.Test.APIKey, setup.StorageTypeAirTable.Test.TableName}

	// pass to dedicated memory tester, to test against storager interface
	tester := Tester{&s, t}
	tester.Test()
}
