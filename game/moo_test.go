package game

import "testing"

func TestGetHit(t *testing.T) {
	{
		g := []int{1, 2, 3, 4}
		a := []int{8, 2, 4, 3}
		expected := 1
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{1, 2, 3, 4}
		expected := 4
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{5, 6, 7, 9}
		expected := 0
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
}

func TestGetBlow(t *testing.T) {
	{
		g := []int{1, 2, 3, 4}
		a := []int{8, 2, 4, 3}
		expected := 2
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{1, 2, 3, 4}
		expected := 0
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{4, 3, 2, 1}
		expected := 4
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
}
