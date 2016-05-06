package pod

import (
	"testing"
	"fmt"
	"strconv"
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

func get_p4() []*Pod {
	nodes := map[string]map[string]int{}
	var result []*Pod
	n1 := map[string]int {
		"0":10, "1":10, "2":10, "3":10, "4":10, "5":10,
	}
	n2 := map[string]int {
		"0":10, "1":10, "2":10, "3":10, "4":10, "5":10,
	}
	n31 := map[string]int {
		"0":10, "1":10,
	}
	n32 := map[string]int {
		"0":10, "1":10, "2":10,
	}
	n33 := map[string]int {
		"0":10, "1":10, "2":10, "3":10,
	}
	nodes["n1"] = n1
	nodes["n2"] = n2
	result = append(result, NewPod(nodes, -1, 10))
	for i := 0; i < 3; i++ {
		if i == 0 {
			nodes["n3"] = n31
		}
		if i == 1 {
			nodes["n3"] = n32
		}
		if i == 2 {
			nodes["n3"] = n33
		}
	}
	result = append(result, NewPod(nodes, -1, 10))
	return result
}

func TestNotEvenAvgPlan(t *testing.T) {
	ps := get_p4()
	for k, p := range ps {
		for i := 0; i < 10; i++ {
			res := p.AveragePlan(1.3, 10, 4)
			if k == 2 && res != nil {
				fmt.Println(p.nodes)
				t.Fatalf("Something went wrong")
			} else if res == nil {
				t.Fatalf("Not enough resource")
			}
		}
	}
}

func TestEvenAvgPlan(t *testing.T) {
	ps := get_p4()
	for _, p := range ps {
		for i := 0; i < 10; i++ {
			res := p.AveragePlan(1.3, 8, 4)
			if res == nil {
				t.Fatalf("not enough resource")
			}
		}
	}
}

func genEvenPod(hostNum, coreNum, coreShare, maxShareCore int) *Pod {
	nodes := map[string]map[string]int{}
	for i := 0; i < hostNum; i++ {
		key := "host" + strconv.Itoa(i)
		nodes[key] = map[string]int{}
		for j := 0; j < coreNum; j++ {
			nodes[key][strconv.Itoa(j)] = coreShare
		}
	}
	pod := NewPod(nodes, maxShareCore, coreShare)
	return pod
}

func TestCenterPlan(t *testing.T) {
	pod1 := genEvenPod(5, 10, 10, -1)
	pod2 := genEvenPod(3, 5, 10, -1)

	res1 := pod1.CenterPlan(1.3, 35)
	if res1 == nil {
		t.Fatalf("test1: not enough resource")
	}
	res2 := pod1.CenterPlan(1.7, 25)
	if res2 == nil {
		t.Fatalf("test2: not enough resource")
	}

	res3 := pod1.AveragePlan(1.3, 30, 7)
	if res3 == nil {
		t.Fatalf("test3: not enough resource")
	}
	res4 := pod1.AveragePlan(1.7, 21, 5)
	if res4 == nil {
		t.Fatalf("test4: not enough resource")
	}

	res5 := pod2.CenterPlan(1.3, 18)
	if res5 != nil {
		t.Fatalf("test5: sth went wrong")
	}
	res6 := pod2.CenterPlan(1.7, 13)
	if res6 != nil {
		t.Fatalf("test6: sth went wrong")
	}

	res7 := pod2.AveragePlan(1.3, 10, 3)
	if res7 != nil {
		fmt.Println(res7)
		t.Fatalf("test7: sth went wrong")
	}
	res8 := pod2.AveragePlan(1.7, 7, 2)
	if res8 != nil {
		fmt.Println(res8)
		t.Fatalf("test8: sth went wrong")
	}

}
