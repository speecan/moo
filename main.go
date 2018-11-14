package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/speecan/moo/game"
	"github.com/speecan/moo/sample"
)

var (
	difficulty = 4    // moo digit number <= 10
	benchNum   = 1    // how many benchmark
	workers    = 16   // run workers at the same time
	queueSize  = 1000 // queue size for worker depend on PC memoty size

	totalEstimates = 0
	totalDuration  = time.Duration(0)

	wg    = &sync.WaitGroup{}
	mutex = &sync.Mutex{}
)

func main() {
	// init
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	game.DebugMode = false
	queue := make(chan game.Estimate, queueSize)
	wg.Add(workers)

	// run benchmark for moo estimater
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for fn := range queue {
				startTime := time.Now()
				count := runOnce(fn) // 1回数当てゲームを行う
				duration := time.Since(startTime)

				// かかった時間の集計
				mutex.Lock()
				totalEstimates += count
				totalDuration += duration
				mutex.Unlock()
			}
		}()
	}

	for n := 0; n < benchNum; n++ {
		queue <- sample.EstimateHuman(difficulty)
		// queue <- sample.EstimateWithRandom(difficulty)
		// queue <- sample.EstimateWithRandom2(difficulty)
		// queue <- munenari.Estimate(difficulty)
	}
	close(queue)
	wg.Wait()

	// result
	fmt.Println("avg. spent:", totalDuration/time.Duration(benchNum), "avg. estimates count:", float64(totalEstimates)/float64(benchNum))
}

func runOnce(estimateFn game.Estimate) int {
	// set game difficulty to 4
	g := game.NewGame(difficulty)
	count := 0
	q := g.GetQuestion(&count)
	fmt.Println("answer:", g.GetAnswer())
	for {
		// loop until hit the answer
		res := estimateFn(q)
		if g.Equals(res) {
			break
		}
	}
	fmt.Println("total questions:", count)
	return count
}
