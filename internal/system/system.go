package system

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ozanbozkurt/hardware-benchmark/internal/types"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSystemInfo() types.SystemInfo {
	info := types.SystemInfo{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		CPUCount:     runtime.NumCPU(),
	}

	// Get detailed CPU info
	cpuInfo, _ := cpu.Info()
	if len(cpuInfo) > 0 {
		info.CPUModel = cpuInfo[0].ModelName
		info.CPUCores = int(cpuInfo[0].Cores)
		info.CPUThreads = int(cpuInfo[0].Cores)
	}

	// Get memory info
	memStat, _ := mem.VirtualMemory()
	info.TotalMemory = memStat.Total

	// Get hostname
	if hostname, err := os.Hostname(); err == nil {
		info.Hostname = hostname
	}

	// Fix CPU count if needed
	if info.CPUCount == 0 {
		info.CPUCount = 1
	}
	if info.CPUCores == 0 {
		info.CPUCores = info.CPUCount
	}
	if info.CPUThreads == 0 {
		info.CPUThreads = info.CPUCount
	}

	return info
}

func GetCPUModel() string {
	info, err := host.Info()
	if err == nil {
		return fmt.Sprintf("%s", info.OS)
	}
	return "Unknown"
}
