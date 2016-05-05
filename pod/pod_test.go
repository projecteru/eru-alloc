package pod

import (
	"testing"
)

func get_p() *Pod {
	nodes := map[string]map[string]int{}
	n1 := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	n2 := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	n3 := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	n4 := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	n5 := map[string]int{
		"0": 10, "1": 10, "2": 10, "3": 10,
		"4": 10, "5": 10, "6": 10, "7": 10,
		"8": 10, "9": 10, "10": 10, "11": 10,
	}
	nodes["n1"] = n1
	nodes["n2"] = n2
	nodes["n3"] = n3
	nodes["n4"] = n4
	nodes["n5"] = n5
	p := NewPod(nodes, -1, 10)
	return p
}

func Benchmark_Center(b *testing.B) {
	b.StopTimer()
	p := get_p()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.CenterPlan(1.3, 44)
	}
}

func Benchmark_Avg(b *testing.B) {
	b.StopTimer()
	p := get_p()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.AveragePlan(1.3, 9, 2)
	}
}
