package main

import (
	"sync"
	"time"
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

func main() {}
