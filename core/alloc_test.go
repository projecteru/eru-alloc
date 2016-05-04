package core

import "testing"

func Benchmark_Alloc(b *testing.B) {
	b.StopTimer()
	cpuInfo := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		host := NewHost(cpuInfo, 10)
		host.GetContainerCores(1.1, -1)
	}
}
