# moo

## setting

### Edit benchNum in main.go (recommend: 10000)

already set 1 for human input/estimating

### fill Estimate func type

以下の関数を実装し、 `main.go` から呼び出す

```go
func(difficulty int) game.Estimate
```

`game.Estimate` の func は1度のゲーム(answerが同じ間)の中で何度も call される

`game.Question` を予想した答えと共に実行すると、これの `Hit` と `Blow` が返却され質問した回数がカウントされる

実装した `func(difficulty int) game.Estimate` は `goroutine` セーフである必要がある

> ベンチマークのため、並列で呼び出される

`game` package は自由に利用して良い

## run benchmark

```bash
$ go run main.go
```

- example.)

```
~~~
avg. spent: 883.156µs avg. estimates count: 5.5789
```

> 理論値不明 (5.3くらい?)

for using an example algorithm (too idiot)

`sample/dummy.go`

```go
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}
```

`main.go`

```go
	for n := 0; n < benchNum; n++ {
		queue <- sample.EstimateHuman(difficulty)
	}
```
