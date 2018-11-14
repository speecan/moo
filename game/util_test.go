package game

import (
	"sort"
	"testing"
)

func TestEquals(t *testing.T) {
	{
		a := []int{1, 2, 3, 4, 5, 6, 7}
		b := []int{1, 2, 3, 4, 5, 6, 7}
		res := Equals(a, b)
		if !res {
			t.Error(a, "and", b, "must be equals")
		}
	}
	{
		a := []int{1, 2, 3, 4, 5, 6, 7}
		b := []int{9, 2, 3, 4, 5, 6, 7}
		res := Equals(a, b)
		if res {
			t.Error(a, "and", b, "must be not equals")
		}
	}
}

func TestStr2Int(t *testing.T) {
	{
		a := []string{"0", "1", "8", "5"}
		expected := []int{0, 1, 8, 5}
		res := Str2Int(a)
		if !Equals(res, expected) {
			t.Error(a, "converted", res, "expected:", expected)
		}
	}
	{
		a := []string{"0", "1", "8", "g", "5"}
		// string was not number will be ignored
		expected := []int{0, 1, 8, 5}
		res := Str2Int(a)
		if !Equals(res, expected) {
			t.Error(a, "converted", res, "expected:", expected)
		}
	}
}

func TestShuffle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	b := []int{0, 1, 2, 3, 4, 5}
	Shuffle(a)
	if len(a) != 6 {
		t.Error("slice should be just shuffled")
	}
	sort.Ints(a)
	if !Equals(a, b) {
		t.Error(a, b, "shuffle and sort then it should be same")
	}
}

func TestGetMooNum(t *testing.T) {
	res := GetMooNum(4)
	if len(res) != 4 {
		t.Error(res, "must be, ", 4, "length")
	}
}
