package memory

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/ozanbozkurt/hardware-benchmark/internal/types"
)

func RunMemoryBenchmark() types.MemoryResult {
	result := types.MemoryResult{}

	// Sequential read/write bandwidth
	result.ReadBandwidth, result.WriteBandwidth = bandwidthBenchmark()

	// Memory latency
	result.Latency = latencyBenchmark()

	// Health score
	result.HealthScore = getMemoryHealth()

	return result
}

func bandwidthBenchmark() (float64, float64) {
	bufferSize := 100 * 1024 * 1024 // 100MB
	buffer := make([]byte, bufferSize)
	iterations := 10

	// Read bandwidth
	start := time.Now()
	for iter := 0; iter < iterations; iter++ {
		sum := 0
		for i := 0; i < len(buffer); i += 64 {
			sum += int(buffer[i])
		}
		_ = sum
	}
	readTime := time.Since(start)
	readBandwidth := float64(bufferSize*iterations) / (1024 * 1024) / readTime.Seconds()

	// Write bandwidth
	start = time.Now()
	for iter := 0; iter < iterations; iter++ {
		for i := 0; i < len(buffer); i += 64 {
			buffer[i] = byte(iter)
		}
	}
	writeTime := time.Since(start)
	writeBandwidth := float64(bufferSize*iterations) / (1024 * 1024) / writeTime.Seconds()

	return readBandwidth, writeBandwidth
}

func latencyBenchmark() float64 {
	// Array access latency test
	arraySize := 256 * 1024 // 256KB (typically fits in L3 cache)
	array := make([]uint32, arraySize)

	// Pseudo-random access pattern
	index := 0
	start := time.Now()
	iterations := 1000000

	for i := 0; i < iterations; i++ {
		index = (index + 167) % arraySize
		_ = array[index]
	}

	duration := time.Since(start)
	latency := duration.Nanoseconds() / int64(iterations)

	if latency < 1 {
		latency = 1
	}

	return float64(latency)
}

func getMemoryHealth() float64 {
	// Test memory by writing and reading patterns
	testSize := 10 * 1024 * 1024 // 10MB
	buffer := make([]byte, testSize)
	healthScore := 100.0

	// Test 1: Sequential write/read
	pattern := byte(0xAA)
	for i := 0; i < testSize; i++ {
		buffer[i] = pattern
	}

	errors := 0
	for i := 0; i < testSize; i++ {
		if buffer[i] != pattern {
			errors++
		}
	}

	if errors > 0 {
		healthScore -= float64(errors) / float64(testSize) * 20
	}

	// Test 2: Alternating pattern
	pattern = 0x55
	for i := 0; i < testSize; i++ {
		buffer[i] = pattern
	}

	errors = 0
	for i := 0; i < testSize; i++ {
		if buffer[i] != pattern {
			errors++
		}
	}

	if errors > 0 {
		healthScore -= float64(errors) / float64(testSize) * 20
	}

	// Test 3: Random pattern stress
	for i := 0; i < testSize; i++ {
		buffer[i] = byte((i * 7) % 256)
	}

	errors = 0
	for i := 0; i < testSize; i++ {
		expected := byte((i * 7) % 256)
		if buffer[i] != expected {
			errors++
		}
	}

	if errors > 0 {
		healthScore -= float64(errors) / float64(testSize) * 20
	}

	// Prevent negative scores
	if healthScore < 0 {
		healthScore = 0
	}

	// Factor in memory usage pressure
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memPressure := float64(m.Alloc) / float64(m.TotalAlloc)
	if memPressure > 0.9 {
		healthScore -= 10
	}

	// Ensure reasonable bounds
	if healthScore > 100 {
		healthScore = 100
	}
	if healthScore < 0 {
		healthScore = 0
	}

	return healthScore
}

// Pointer to ensure memory operations aren't optimized away
func volatileRead(ptr *byte) byte {
	return *ptr
}

func volatileWrite(ptr *byte, val byte) {
	*ptr = val
	// Force visibility
	_ = unsafe.Pointer(ptr)
}
