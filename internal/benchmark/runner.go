package benchmark

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/ozanbozkurt/hardware-benchmark/internal/cpu"
	"github.com/ozanbozkurt/hardware-benchmark/internal/disk"
	"github.com/ozanbozkurt/hardware-benchmark/internal/memory"
	"github.com/ozanbozkurt/hardware-benchmark/internal/network"
	"github.com/ozanbozkurt/hardware-benchmark/internal/system"
	"github.com/ozanbozkurt/hardware-benchmark/internal/types"
)

type BenchmarkRunner struct {
	startTime time.Time
}

func NewBenchmarkRunner() *BenchmarkRunner {
	return &BenchmarkRunner{
		startTime: time.Now(),
	}
}

func (br *BenchmarkRunner) RunAllBenchmarks() *types.BenchmarkResults {
	results := &types.BenchmarkResults{
		Timestamp: time.Now(),
	}

	fmt.Println("📊 Collecting system information...")
	results.SystemInfo = system.GetSystemInfo()

	fmt.Println("⚙️  Running CPU benchmarks...")
	results.CPU = cpu.RunCPUBenchmark()

	fmt.Println("🧠 Running memory benchmarks...")
	results.Memory = memory.RunMemoryBenchmark()

	fmt.Println("💾 Running disk benchmarks...")
	results.Disk = disk.RunDiskBenchmark()

	fmt.Println("🌐 Running network benchmarks...")
	results.Network = network.RunNetworkBenchmark()

	fmt.Println("🎯 Calculating quality score...")
	results.QualityScore = calculateQualityScore(results)

	return results
}

func calculateQualityScore(results *types.BenchmarkResults) types.QualityScore {
	score := types.QualityScore{
		Timestamp: time.Now(),
	}

	// CPU Score (0-25)
	cpuPercentage := (results.CPU.SingleThreadScore / 1000.0) * 100
	if cpuPercentage > 100 {
		cpuPercentage = 100
	}
	score.CPUScore = cpuPercentage * 0.25

	// Memory Score (0-25)
	memoryBandwidth := (results.Memory.ReadBandwidth / 20000.0) * 100
	if memoryBandwidth > 100 {
		memoryBandwidth = 100
	}
	score.MemoryScore = memoryBandwidth * 0.25

	// Disk Score (0-25)
	diskSpeed := (results.Disk.WriteSpeed / 500.0) * 100
	if diskSpeed > 100 {
		diskSpeed = 100
	}
	score.DiskScore = diskSpeed * 0.25

	// Network Score (0-25)
	networkSpeed := (results.Network.DownloadSpeed / 1000.0) * 100
	if networkSpeed > 100 {
		networkSpeed = 100
	}
	score.NetworkScore = networkSpeed * 0.25

	// Total Score
	score.TotalScore = score.CPUScore + score.MemoryScore + score.DiskScore + score.NetworkScore

	// Reliability metrics
	score.CPUStability = results.CPU.Stability
	score.MemoryHealth = results.Memory.HealthScore
	score.DiskHealth = results.Disk.HealthScore

	return score
}

func (br *BenchmarkRunner) PrintResults(results *types.BenchmarkResults) {
	elapsed := time.Since(br.startTime).Seconds()

	fmt.Println("\n" + color.GreenString("╔════════════════════════════════════════════════════════════════╗"))
	fmt.Println(color.GreenString("║") + "          HARDWARE PERFORMANCE BENCHMARK RESULTS                   " + color.GreenString("║"))
	fmt.Println(color.GreenString("╚════════════════════════════════════════════════════════════════╝"))

	// System Info
	fmt.Println("\n" + color.CyanString("📱 SYSTEM INFORMATION"))
	fmt.Println(color.WhiteString("  OS:              ") + results.SystemInfo.OS)
	fmt.Println(color.WhiteString("  Architecture:    ") + results.SystemInfo.Architecture)
	fmt.Println(color.WhiteString("  CPU Count:       ") + fmt.Sprintf("%d (cores: %d, threads: %d)", results.SystemInfo.CPUCount, results.SystemInfo.CPUCores, results.SystemInfo.CPUThreads))
	fmt.Println(color.WhiteString("  Total Memory:    ") + fmt.Sprintf("%.2f GB", float64(results.SystemInfo.TotalMemory)/1024/1024/1024))
	fmt.Println(color.WhiteString("  CPU Model:       ") + results.SystemInfo.CPUModel)
	fmt.Println(color.WhiteString("  Hostname:        ") + results.SystemInfo.Hostname)

	// CPU Results
	fmt.Println("\n" + color.CyanString("⚙️  CPU BENCHMARK"))
	fmt.Println(color.WhiteString("  Single Thread:   ") + color.YellowString(fmt.Sprintf("%.2f pts", results.CPU.SingleThreadScore)))
	fmt.Println(color.WhiteString("  Multi Thread:    ") + color.YellowString(fmt.Sprintf("%.2f pts", results.CPU.MultiThreadScore)))
	fmt.Println(color.WhiteString("  Ops/Sec:         ") + fmt.Sprintf("%.0f", results.CPU.OpsPerSecond))
	fmt.Println(color.WhiteString("  CPU Usage Avg:   ") + fmt.Sprintf("%.2f%%", results.CPU.AverageUsage))
	fmt.Println(color.WhiteString("  Stability:       ") + getStabilityColor(results.CPU.Stability))

	// Memory Results
	fmt.Println("\n" + color.CyanString("🧠 MEMORY BENCHMARK"))
	fmt.Println(color.WhiteString("  Read Bandwidth:  ") + color.YellowString(fmt.Sprintf("%.2f MB/s", results.Memory.ReadBandwidth)))
	fmt.Println(color.WhiteString("  Write Bandwidth: ") + color.YellowString(fmt.Sprintf("%.2f MB/s", results.Memory.WriteBandwidth)))
	fmt.Println(color.WhiteString("  Latency:         ") + fmt.Sprintf("%.2f ns", results.Memory.Latency))
	fmt.Println(color.WhiteString("  Health Score:    ") + fmt.Sprintf("%.2f/100", results.Memory.HealthScore))

	// Disk Results
	fmt.Println("\n" + color.CyanString("💾 DISK BENCHMARK"))
	fmt.Println(color.WhiteString("  Read Speed:      ") + color.YellowString(fmt.Sprintf("%.2f MB/s", results.Disk.ReadSpeed)))
	fmt.Println(color.WhiteString("  Write Speed:     ") + color.YellowString(fmt.Sprintf("%.2f MB/s", results.Disk.WriteSpeed)))
	fmt.Println(color.WhiteString("  4K Read IOPS:    ") + fmt.Sprintf("%.0f", results.Disk.RandomReadIOPS))
	fmt.Println(color.WhiteString("  4K Write IOPS:   ") + fmt.Sprintf("%.0f", results.Disk.RandomWriteIOPS))
	fmt.Println(color.WhiteString("  Health Score:    ") + fmt.Sprintf("%.2f/100", results.Disk.HealthScore))

	// Network Results
	fmt.Println("\n" + color.CyanString("🌐 NETWORK BENCHMARK"))
	fmt.Println(color.WhiteString("  Download:        ") + color.YellowString(fmt.Sprintf("%.2f Mbps", results.Network.DownloadSpeed)))
	fmt.Println(color.WhiteString("  Upload:          ") + color.YellowString(fmt.Sprintf("%.2f Mbps", results.Network.UploadSpeed)))
	fmt.Println(color.WhiteString("  Ping (ms):       ") + fmt.Sprintf("%.2f", results.Network.Latency))
	fmt.Println(color.WhiteString("  Packet Loss:     ") + fmt.Sprintf("%.2f%%", results.Network.PacketLoss))

	// Quality Score
	fmt.Println("\n" + color.GreenString("🎯 OVERALL QUALITY SCORE"))
	printQualityScore(results.QualityScore)

	// Recommendations
	fmt.Println("\n" + color.MagentaString("💡 RECOMMENDATIONS"))
	printRecommendations(results)

	fmt.Printf("\n⏱️  Total benchmark time: %.2f seconds\n\n", elapsed)
}

func getStabilityColor(stability float64) string {
	if stability >= 95 {
		return color.GreenString(fmt.Sprintf("%.2f%% (Excellent)", stability))
	} else if stability >= 85 {
		return color.YellowString(fmt.Sprintf("%.2f%% (Good)", stability))
	} else if stability >= 70 {
		return color.HiYellowString(fmt.Sprintf("%.2f%% (Fair)", stability))
	}
	return color.RedString(fmt.Sprintf("%.2f%% (Poor)", stability))
}

func printQualityScore(qs types.QualityScore) {
	total := qs.TotalScore

	var grade string
	var colorFunc func(string, ...interface{}) string

	switch {
	case total >= 90:
		grade = "A+ (Premium)"
		colorFunc = color.GreenString
	case total >= 80:
		grade = "A (Excellent)"
		colorFunc = color.GreenString
	case total >= 70:
		grade = "B (Very Good)"
		colorFunc = color.YellowString
	case total >= 60:
		grade = "C (Good)"
		colorFunc = color.YellowString
	case total >= 50:
		grade = "D (Fair)"
		colorFunc = color.HiYellowString
	default:
		grade = "F (Poor)"
		colorFunc = color.RedString
	}

	fmt.Printf("  %s Total Score: %s (%.2f/100)\n",
		"├",
		colorFunc(grade),
		total,
	)
	fmt.Printf("  %s CPU:         %.2f/25\n", "├", qs.CPUScore)
	fmt.Printf("  %s Memory:      %.2f/25\n", "├", qs.MemoryScore)
	fmt.Printf("  %s Disk:        %.2f/25\n", "├", qs.DiskScore)
	fmt.Printf("  %s Network:     %.2f/25\n", "└", qs.NetworkScore)
}

func printRecommendations(results *types.BenchmarkResults) {
	qs := results.QualityScore

	if qs.CPUScore < 15 {
		fmt.Println("  ⚠️  CPU performance is below expected. Consider upgrading.")
	}
	if qs.MemoryScore < 15 {
		fmt.Println("  ⚠️  Memory bandwidth is limited. May need RAM upgrade.")
	}
	if qs.DiskScore < 15 {
		fmt.Println("  ⚠️  Disk speed is slow. Consider SSD upgrade or check disk health.")
	}
	if results.Disk.HealthScore < 70 {
		fmt.Println("  🚨 Disk health is concerning. Backup data immediately!")
	}
	if results.CPU.Stability < 80 {
		fmt.Println("  ⚠️  CPU thermal stability is unstable. Check cooling system.")
	}
	if results.Memory.HealthScore < 70 {
		fmt.Println("  ⚠️  Memory may have issues. Run diagnostics.")
	}

	if qs.TotalScore >= 80 {
		fmt.Println("  ✅ System is in excellent condition!")
	} else if qs.TotalScore >= 60 {
		fmt.Println("  ✅ System performs well for most tasks.")
	}
}
