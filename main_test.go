package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/speecan/moo/game"
	"github.com/speecan/moo/sample"
)

func TestRunOnce(t *testing.T) {
	count := runOnce(sample.EstimateWithRandom(difficulty))
	if count == 0 {
		t.Error("estimate may be invalid")
	}
	// t.Error("for debug")
}

func BenchmarkRunOnceParallel(b *testing.B) {
	game.DebugMode = false
	mutex := &sync.Mutex{}
	count := 0
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c := runOnce(sample.EstimateWithRandom(difficulty))
			// c := runOnce(sample.EstimateWithRandom(difficulty))
			mutex.Lock()
			count += c
			mutex.Unlock()
		}
	})
	fmt.Println("avg. estimates count", float64(count)/float64(b.N))
}
