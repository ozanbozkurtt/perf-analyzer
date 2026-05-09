package network

import (
	"net"
	"time"

	"github.com/ozanbozkurt/hardware-benchmark/internal/types"
)

func RunNetworkBenchmark() types.NetworkResult {
	result := types.NetworkResult{}

	// Network connectivity test
	result.DownloadSpeed = estimateNetworkSpeed()
	result.UploadSpeed = result.DownloadSpeed * 0.8 // Typical asymmetric ratio

	// Latency test
	result.Latency = measureLatency()

	// Packet loss estimation
	result.PacketLoss = estimatePacketLoss()

	return result
}

func estimateNetworkSpeed() float64 {
	// Simple bandwidth estimation using DNS lookups
	testHosts := []string{
		"8.8.8.8:53",         // Google DNS
		"1.1.1.1:53",         // Cloudflare DNS
		"208.67.222.222:53",  // OpenDNS
		"127.0.0.1:22",       // localhost fallback
	}

	totalLatency := 0.0
	successCount := 0

	for _, host := range testHosts {
		if latency := testConnectivity(host); latency > 0 {
			totalLatency += latency
			successCount++
		}
	}

	if successCount == 0 {
		return 5.0 // Fallback: assume slow connection
	}

	avgLatency := totalLatency / float64(successCount)

	// Estimate based on latency (rough estimate)
	// Lower latency = higher potential bandwidth
	if avgLatency < 10 {
		return 100.0 // Good connection
	} else if avgLatency < 50 {
		return 50.0 // Moderate connection
	} else if avgLatency < 150 {
		return 20.0 // Slow connection
	}
	return 5.0 // Very slow connection
}

func testConnectivity(address string) float64 {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return 0.0
	}
	defer conn.Close()

	return float64(time.Since(start).Milliseconds())
}

func measureLatency() float64 {
	// Test latency to multiple targets
	targets := []string{
		"8.8.8.8:53",
		"1.1.1.1:53",
		"208.67.222.222:53",
	}

	totalLatency := 0.0
	count := 0

	for _, target := range targets {
		latency := pingHost(target)
		if latency > 0 {
			totalLatency += latency
			count++
		}
	}

	if count == 0 {
		return 50.0 // Default if no connectivity
	}

	return totalLatency / float64(count)
}

func pingHost(address string) float64 {
	attempts := 3
	totalTime := 0.0

	for i := 0; i < attempts; i++ {
		start := time.Now()
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		totalTime += float64(time.Since(start).Milliseconds())
	}

	if totalTime == 0 {
		return 0.0
	}

	return totalTime / float64(attempts)
}

func estimatePacketLoss() float64 {
	// Estimate packet loss based on connectivity reliability
	targets := []string{
		"8.8.8.8:53",
		"1.1.1.1:53",
		"208.67.222.222:53",
	}

	failureCount := 0
	totalAttempts := len(targets) * 3

	for _, target := range targets {
		for i := 0; i < 3; i++ {
			if testConnectivity(target) == 0 {
				failureCount++
			}
		}
	}

	if totalAttempts == 0 {
		return 0.0
	}

	return (float64(failureCount) / float64(totalAttempts)) * 100.0
}
