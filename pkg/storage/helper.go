package storage

import "sort"

// ReverseNotesArray reverse order given notes array
func ReverseNotesArray(list []Note) []Note {
	var reversed []Note
	count := len(list)

	for i := count; i > 0; i-- {
		reversed = append(reversed, list[i-1])
	}
	return reversed
}

// SortByIDDesc sorts given notes list by ID desc
func SortByIDDesc(list []Note) []Note {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID > list[j].ID
	})
	return list
}
