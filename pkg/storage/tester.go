package storage

import (
	"fmt"
	"testing"
)

// Tester structure dedicated to test different storage implementation agains storager interface
type Tester struct {
	Storage Storager
	T       *testing.T
}

func (s *Tester) Test() {
	s.testDeleteAll()

	firstContent := "First"
	secondContent := "Second Note"
	thirdContent := "Third note haha"

	firstID := s.testAdd(firstContent)
	secondID := s.testAdd(secondContent)
	thirdID := s.testAdd(thirdContent)

	s.testCount(3)

	s.compareNotes(
		[]Note{Note{thirdID, thirdContent, ""}, Note{secondID, secondContent, ""}, Note{firstID, firstContent, ""}},
		s.testGet(),
		"Get. Comparing notes after 3 inserts")

	s.compareNotes(
		[]Note{Note{thirdID, thirdContent, ""}, Note{secondID, secondContent, ""}, Note{firstID, firstContent, ""}},
		s.testGetN(4),
		"GetN(4) Comparing notes after 3 inserts")

	s.compareNotes(
		[]Note{Note{thirdID, thirdContent, ""}, Note{secondID, secondContent, ""}},
		s.testGetN(2),
		"GetN(2) Comparing notes after 3 inserts")

	s.testDelete(secondID)

	s.compareNotes(
		[]Note{Note{thirdID, thirdContent, ""}, Note{firstID, firstContent, ""}},
		s.testGet(),
		"Get() After delete")

	s.testDeleteAll()
}

func (s *Tester) testDeleteAll() {
	err := s.Storage.DeleteAll()
	if err != nil {
		s.logError("DeleteAll error: " + err.Error())
	}

	count := s.countItems()

	if count > 0 {
		s.logError(fmt.Sprintf("After 'DeleteAll' There is still %v notes in storage.", count))
	}
}

func (s *Tester) testDelete(ID int) {
	err := s.Storage.Delete(ID)
	if err != nil {
		s.logError(fmt.Sprintf("Delete(%v) failed: ", ID) + err.Error())
	}
}

func (s *Tester) countItems() int {
	list, err := s.Storage.Get()
	if err != nil {
		s.logError("Get error: " + err.Error())
	}

	return len(list)
}

func (s *Tester) testAdd(content string) (id int) {
	total := s.countItems()
	id, err := s.Storage.Add(content)

	if err != nil {
		s.logError("Add errror: " + err.Error())
	}

	newTotal := s.countItems()
	if newTotal != total+1 {
		s.logError("After new item added quantity not increased by one")
	}

	return
}

func (s *Tester) testCount(expected int) {
	actual := s.countItems()
	if expected != actual {
		s.logError(fmt.Sprintf("Count does not match expected %v, actual %v", expected, actual))
	}
}

func (s *Tester) logError(err string) {
	s.T.Fatal(fmt.Sprintf("Storage type '%T': ", s.Storage) + err)
}

func (s *Tester) testGet() []Note {
	list, err := s.Storage.Get()
	if err != nil {
		s.logError("Get errror: " + err.Error())
	}

	return list
}

func (s *Tester) testGetN(quantity int) []Note {
	list, err := s.Storage.GetN(quantity)
	if err != nil {
		s.logError("GetN errror: " + err.Error())
	}

	return list
}

// compare two array of notes (only ids and content, not date)
func (s *Tester) compareNotes(expected []Note, actual []Note, remark string) {
	if len(expected) != len(actual) {
		s.logError("compareNotes error. Remark(" + remark + ") Array not same size")
	}

	for i, note := range actual {
		if note.ID != expected[i].ID || note.Content != expected[i].Content {
			s.logError("compareNotes error. Remark(" + remark + "). Notes does not match. Expected: " + fmt.Sprint(expected) + " .Actual: " + fmt.Sprint(actual))
		}
	}
}
