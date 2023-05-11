package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/speecan/moo/game"
	"github.com/speecan/moo/sample"
)

func TestRunOnce(t *testing.T) {
	count := game.Run(difficulty, sample.Random{})
	if count == 0 {
		t.Error("estimate may be invalid")
	}
	// t.Error("for debug")
}

func BenchmarkRunOnceParallel(b *testing.B) {
	game.DebugMode = false
	mutex := &sync.Mutex{}
	count := 0
	e := sample.Random{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c := game.Run(difficulty, e)
			mutex.Lock()
			count += c
			mutex.Unlock()
		}
	})
	fmt.Println("avg. estimates count", float64(count)/float64(b.N))
}

func BenchmarkRun(b *testing.B) {
	game.DebugMode = false
	es := []game.Estimater{
		sample.Random{},
		sample.Random2{},
	}
	for _, e := range es {
		ee := e
		b.Run(fmt.Sprintf("%+v", ee), func(b *testing.B) {
			count := 0
			for i := 0; i < b.N; i++ {
				count += game.Run(difficulty, ee)
			}
			b.Log("avg. estimates count", float64(count)/float64(b.N))
		})
	}
}
