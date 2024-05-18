package gslice

import (
	"errors"
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	arr := []int{1, 2, 3}
	var sum int
	ForEach(arr, func(v int) {
		sum += v
	})
	if sum != 6 {
		t.Errorf("Expected sum to be 6, got %d", sum)
	}
}

func TestTryForEach(t *testing.T) {
	arr := []int{1, 2, 3}
	err := TryForEach(arr, func(v int) error {
		if v == 2 {
			return errors.New("error on 2")
		}
		return nil
	})
	if err == nil || err.Error() != "error on 2" {
		t.Errorf("Expected error 'error on 2', got %v", err)
	}
}

func TestForEach2(t *testing.T) {
	arr := []string{"a", "b", "c"}
	var res string
	ForEach2(arr, func(i int, v string) {
		res += fmt.Sprintf("%d:%s", i, v)
	})
	expected := "0:a1:b2:c"
	if res != expected {
		t.Errorf("Expected %q, got %q", expected, res)
	}
}

func TestTryForEach2(t *testing.T) {
	arr := []int{1, 2, 3}
	err := TryForEach2(arr, func(i int, v int) error {
		if i == 1 {
			return errors.New("error on index 1")
		}
		return nil
	})
	if err == nil || err.Error() != "error on index 1" {
		t.Errorf("Expected error 'error on index 1', got %v", err)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	filtered := Filter(arr, func(v int) bool {
		return v%2 == 0
	})
	expected := []int{2, 4}
	if !slicesEqual(filtered, expected) {
		t.Errorf("Expected %v, got %v", expected, filtered)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	mapped := Map(arr, func(v int) string {
		return fmt.Sprintf("%d", v*2)
	})
	expected := []string{"2", "4", "6"}
	if !slicesEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, mapped)
	}
}

func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
