package game

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	initializedSeed = false
)

// Shuffle slice
func Shuffle(a []int) {
	if !initializedSeed {
		rand.Seed(time.Now().UnixNano())
	}
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Equals compares guess and answer that are equal
func Equals(guess, answer []int) bool {
	if len(guess) != len(answer) {
		return false
	}
	for i, g := range guess {
		if answer[i] != g {
			return false
		}
	}
	return true
}

// Str2Int convert []string to []int
// string that was not int will be ignored (is not filled in)
func Str2Int(str []string) []int {
	res := make([]int, 0)
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		res = append(res, i)
	}
	return res
}

// GetMooNum returns moo number with digit
func GetMooNum(digit int) []int {
	sl := make([]int, len(nums))
	copy(sl, nums)
	Shuffle(sl)
	answer := sl[:digit]
	return answer
}
