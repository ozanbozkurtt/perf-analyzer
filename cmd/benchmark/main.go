package main

import (
	"fmt"
	"os"

	"github.com/ozanbozkurt/hardware-benchmark/internal/benchmark"
)

func main() {
	fmt.Println("\n🔥 Hardware Performance Benchmark Tool\n")

	runner := benchmark.NewBenchmarkRunner()
	results := runner.RunAllBenchmarks()

	runner.PrintResults(results)

	os.Exit(0)
}
