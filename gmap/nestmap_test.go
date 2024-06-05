package gmap

import "testing"

func TestGetSetNested(t *testing.T) {
	// Test case 1: Testing with a non-nil map
	m := make(map[int]map[string]int)
	k := 1
	v := GetSetNested(m, k)
	if v == nil {
		t.Errorf("Expected non-nil map, got nil")
	}

	// Test case 2: Testing with a nil map
	m = nil
	v = GetSetNested(m, k)
	if v != nil {
		t.Errorf("Expected nil map, got non-nil")
	}
}
