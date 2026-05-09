package cpu

import (
	"sync"
	"time"

	"github.com/ozanbozkurt/hardware-benchmark/internal/types"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

func RunCPUBenchmark() types.CPUResult {
	result := types.CPUResult{}

	// Single thread benchmark
	result.SingleThreadScore = singleThreadBenchmark()

	// Multi thread benchmark
	result.MultiThreadScore = multiThreadBenchmark()

	// Operations per second
	result.OpsPerSecond = opsPerSecondBenchmark()

	// CPU usage during benchmark
	result.AverageUsage = getCPUUsage()

	// Thermal stability
	result.Stability = getThermalStability()

	return result
}

func singleThreadBenchmark() float64 {
	start := time.Now()
	iterations := 0
	result := 0

	for time.Since(start) < time.Second {
		result += complexCalculation(iterations)
		iterations++
	}

	score := float64(iterations) / time.Since(start).Seconds()
	return score * 10
}

func multiThreadBenchmark() float64 {
	numCPU := 4 // Test with 4 goroutines
	start := time.Now()
	results := make(chan int64, numCPU)

	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			count := int64(0)
			for time.Since(start) < time.Second {
				count += int64(simpleOp(id))
			}
			results <- count
		}(i)
	}

	wg.Wait()
	close(results)

	totalOps := int64(0)
	for r := range results {
		totalOps += r
	}

	score := float64(totalOps) / time.Since(start).Seconds()
	return score * 10
}

func opsPerSecondBenchmark() float64 {
	start := time.Now()
	count := 0

	for time.Since(start) < 500*time.Millisecond {
		count += simpleOp(count)
	}

	return float64(count) / time.Since(start).Seconds()
}

func getCPUUsage() float64 {
	p, _ := process.NewProcess(int32(1))
	if p != nil {
		if usage, err := p.CPUPercent(); err == nil {
			return usage
		}
	}

	// Fallback: estimate from system
	if cpuPercent, err := cpu.Percent(500*time.Millisecond, false); err == nil {
		if len(cpuPercent) > 0 {
			return cpuPercent[0]
		}
	}

	return 50.0
}

func getThermalStability() float64 {
	// Measure CPU stability during load
	numGoroutines := 8
	results := make(chan float64, 10)
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func() {
			tick := time.NewTicker(100 * time.Millisecond)
			defer tick.Stop()

			for range tick.C {
				if time.Since(start) > 2*time.Second {
					break
				}
				if cpuPercent, err := cpu.Percent(10*time.Millisecond, false); err == nil {
					if len(cpuPercent) > 0 {
						results <- cpuPercent[0]
					}
				}
			}
		}()
	}

	// Collect measurements
	measurements := []float64{}
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	timeout := time.NewTimer(3 * time.Second)
	defer timeout.Stop()

	for {
		select {
		case <-timeout.C:
			goto analyze
		case <-tick.C:
			select {
			case m := <-results:
				measurements = append(measurements, m)
			default:
			}
		}
	}

analyze:
	if len(measurements) == 0 {
		return 85.0
	}

	// Calculate variance to determine stability
	mean := 0.0
	for _, m := range measurements {
		mean += m
	}
	mean /= float64(len(measurements))

	variance := 0.0
	for _, m := range measurements {
		variance += (m - mean) * (m - mean)
	}
	variance /= float64(len(measurements))

	stdDev := variance
	if stdDev > 0 {
		stdDev = 1.0 / (1.0 + (variance / 100.0)) * 100.0
	} else {
		stdDev = 95.0
	}

	if stdDev > 100 {
		stdDev = 100
	}
	if stdDev < 0 {
		stdDev = 50
	}

	return stdDev
}

func complexCalculation(seed int) int {
	result := seed
	for i := 0; i < 100; i++ {
		result = (result*1103515245 + 12345) & 0x7fffffff
		result = (result * result) / ((seed + 1) * 1000)
	}
	return result
}

func simpleOp(n int) int {
	return (n + 1) * (n + 2) / (n + 1)
}
