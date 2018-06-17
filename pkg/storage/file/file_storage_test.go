package storage

import (
	"testing"

	"github.com/nerijusza/go-notes/pkg/storage"
)

func TestFileStorage(t *testing.T) {
	// inittialize memory storage
	s := FileStorage{}
	err := s.Init("unit_test_interface_tester.txt")
	if err != nil {
		t.Fatal("Object initializiation failed: " + err.Error())
	}

	// pass to dedicated memory tester, to test against storager interface
	tester := storage.Tester{&s, t}
	tester.Test()
}

func TestGetBiggestNoteID(t *testing.T) {
	s := FileStorage{}
	s.Init("unit_test_TestGetBiggestNoteID.json")

	_, err := s.Get()
	if err != nil {
		t.Fatal("First Get failed: " + err.Error())
	}
	return
	err = s.DeleteAll()
	if err != nil {
		t.Fatal("DeleteAll failed: " + err.Error())
	}

	testIfGetBiggestNoteIDWorks(t, s, 0, "Empty file 1")
	testIfGetBiggestNoteIDWorks(t, s, 0, "Empty file 2")

	ID1, err := s.Add("First note")
	if err != nil {
		t.Fatal("Add first note failed: " + err.Error())
	}
	testIfGetBiggestNoteIDWorks(t, s, ID1, "After first insert")

	ID2, err := s.Add("Second note")
	if err != nil {
		t.Fatal("Add second note failed: " + err.Error())
	}
	testIfGetBiggestNoteIDWorks(t, s, ID2, "After second insert")

	err = s.Delete(2)
	if err != nil {
		t.Fatal("Delete second note failed: " + err.Error())
	}

	testIfGetBiggestNoteIDWorks(t, s, ID1, "After delete")

	err = s.DeleteAll()
	if err != nil {
		t.Fatal("DeleteAll failed: " + err.Error())
	}

	testIfGetBiggestNoteIDWorks(t, s, 0, "After DeleteAll")
}

func testIfGetBiggestNoteIDWorks(t *testing.T, s FileStorage, expectedID int, step string) {
	actualID, err := s.getBiggestNoteID()

	if err != nil {
		t.Fatalf("Step (%v). Get ID failed, message: "+err.Error(), step, expectedID)
	}

	if actualID != expectedID {
		t.Fatalf("Step (%v). Expected ID(%v) is not equal actual ID(%v)", step, expectedID, actualID)
	}
}
