package types

import "time"

// SystemInfo holds system information
type SystemInfo struct {
	OS           string
	Architecture string
	CPUCount     int
	CPUCores     int
	CPUThreads   int
	CPUModel     string
	TotalMemory  uint64
	Hostname     string
}

// CPUResult holds CPU benchmark results
type CPUResult struct {
	SingleThreadScore float64
	MultiThreadScore  float64
	OpsPerSecond      float64
	AverageUsage      float64
	Stability         float64
}

// MemoryResult holds memory benchmark results
type MemoryResult struct {
	ReadBandwidth  float64
	WriteBandwidth float64
	Latency        float64
	HealthScore    float64
}

// DiskResult holds disk benchmark results
type DiskResult struct {
	ReadSpeed        float64
	WriteSpeed       float64
	RandomReadIOPS   float64
	RandomWriteIOPS  float64
	HealthScore      float64
}

// NetworkResult holds network benchmark results
type NetworkResult struct {
	DownloadSpeed float64
	UploadSpeed   float64
	Latency       float64
	PacketLoss    float64
}

// QualityScore holds overall quality scoring
type QualityScore struct {
	Timestamp      time.Time
	CPUScore       float64
	MemoryScore    float64
	DiskScore      float64
	NetworkScore   float64
	TotalScore     float64
	CPUStability   float64
	MemoryHealth   float64
	DiskHealth     float64
}

// BenchmarkResults holds all benchmark results
type BenchmarkResults struct {
	Timestamp    time.Time
	SystemInfo   SystemInfo
	CPU          CPUResult
	Memory       MemoryResult
	Disk         DiskResult
	Network      NetworkResult
	QualityScore QualityScore
}
