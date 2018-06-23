package config

import (
	"fmt"
	"testing"
)

func TestSplitNotes(t *testing.T) {
	a := GetConfig()

	fmt.Println(a.StorageType)
	fmt.Println(a.StorageTypeFile.Directory)
	fmt.Println(a.StorageTypeFile.File)
}
