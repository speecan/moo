package sample

import (
	"fmt"
	"os"
	"strings"

	"github.com/speecan/moo/game"
)

type Human struct{}

func (Human) Init(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		var input string
		fmt.Print("?: ")
		fmt.Fscanln(os.Stdin, &input)
		guess := game.Str2Int(strings.Split(input, ""))
		fn(guess)
		return guess
	}
}

type Random struct{}

func (Random) Init(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}

type Random2 struct{}

func (Random2) Init(difficulty int) game.Estimate {
	query := make([][]int, 0)
	isDuplicated := func(i []int) bool {
		for _, v := range query {
			if game.Equals(v, i) {
				return true
			}
		}
		return false
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if !isDuplicated(r) {
				break
			}
		}
		fn(r)
		query = append(query, r)
		return r
	}
}
