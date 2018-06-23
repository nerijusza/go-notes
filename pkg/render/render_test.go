package render

import (
	"reflect"
	"testing"

	"github.com/nerijusza/go-notes/pkg/storage"
)

func TestSplitNotes(t *testing.T) {
	input := []storage.Note{storage.Note{1, "1", "1"}, storage.Note{2, "2", "2"}, storage.Note{3, "3", "3"}, storage.Note{4, "4", "4"}, storage.Note{5, "5", "5"}, storage.Note{6, "6", "6"}, storage.Note{7, "7", "7"}}

	actual := splitNotes(input, 4)
	expected := [][]storage.Note{
		[]storage.Note{storage.Note{1, "1", "1"}, storage.Note{2, "2", "2"}, storage.Note{3, "3", "3"}, storage.Note{4, "4", "4"}},
		[]storage.Note{storage.Note{5, "5", "5"}, storage.Note{6, "6", "6"}, storage.Note{7, "7", "7"}},
	}
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}

	actual = splitNotes(input, 2)
	expected = [][]storage.Note{
		[]storage.Note{storage.Note{1, "1", "1"}, storage.Note{2, "2", "2"}},
		[]storage.Note{storage.Note{3, "3", "3"}, storage.Note{4, "4", "4"}},
		[]storage.Note{storage.Note{5, "5", "5"}, storage.Note{6, "6", "6"}},
		[]storage.Note{storage.Note{7, "7", "7"}},
	}
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}

}
