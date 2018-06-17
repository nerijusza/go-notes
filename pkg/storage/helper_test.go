package storage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseNotesArray(t *testing.T) {
	original := []Note{Note{1, "First", "Date1"}, Note{2, "Second", "Date2"}, Note{3, "Third Note content", "Date3"}}
	expected := []Note{Note{3, "Third Note content", "Date3"}, Note{2, "Second", "Date2"}, Note{1, "First", "Date1"}}

	actual := reverseNotesArray(original)

	if reflect.DeepEqual(expected, actual) == false {
		t.Error(fmt.Sprintf("Error! Expected: %v. Actual: %v", expected, actual))
	}
}

func TestSortByIdDesc(t *testing.T) {
	original := []Note{Note{2, "Second", "Date2"}, Note{1, "First", "Date1"}, Note{4, "Forth", "D4"}, Note{3, "Third Note content", "Date3"}}
	expected := []Note{Note{4, "Forth", "D4"}, Note{3, "Third Note content", "Date3"}, Note{2, "Second", "Date2"}, Note{1, "First", "Date1"}}

	actual := sortByIDDesc(original)

	if reflect.DeepEqual(expected, actual) == false {
		t.Error(fmt.Sprintf("Error! Expected: %v. Actual: %v", expected, actual))
	}
}
