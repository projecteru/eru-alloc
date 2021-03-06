package pod

import "github.com/projecteru/eru-alloc/core"

type Pod struct {
	nodes        map[string]map[string]int
	maxShareCore int
	coreShare    int
}

func NewPod(nodes map[string]map[string]int, maxShareCore, coreShare int) *Pod {
	return &Pod{nodes, maxShareCore, coreShare}
}

func (self *Pod) CenterPlan(cpu float64, need int) map[string][]map[string]int {
	var result = map[string][]map[string]int{}
	var count int = 0

	var n int
	var host *core.Host
	var plan []map[string]int

	for node, cpuInfo := range self.nodes {
		host = core.NewHost(cpuInfo, self.coreShare)
		plan = host.GetContainerCores(cpu, self.maxShareCore)
		n = len(plan)
		if n == 0 {
			continue
		}
		if count+n > need {
			n = need - count
			plan = plan[:n]
		}
		result[node] = plan
		count += n
		if count == need {
			return result
		}
	}
	return nil
}

func (self *Pod) AveragePlan(cpu float64, need int, per int) map[string][]map[string]int {
	var result = map[string][]map[string]int{}
	var host *core.Host
	var plan []map[string]int
	var nodeNum int = need / per
	var last int = need % per
	var n int
	var less = map[string]interface{}{}

	if nodeNum > len(self.nodes) {
		return nil
	}

	for node, cpuInfo := range self.nodes {
		if nodeNum == 0 {
			break
		}
		host = core.NewHost(cpuInfo, self.coreShare)
		plan = host.GetContainerCores(cpu, self.maxShareCore)
		n = len(plan)
		if n >= per {
			result[node] = plan[:per]
			less[node] = new(interface{})
			nodeNum -= 1
		}
	}
	if last == 0 && nodeNum == 0 {
		return result
	} else if last != 0 && nodeNum == 0 {
		for node, cpuInfo := range self.nodes {
			if _, ok := less[node]; !ok {
				host = core.NewHost(cpuInfo, self.coreShare)
				plan = host.GetContainerCores(cpu, self.maxShareCore)
				n = len(plan)
				if n >= last {
					result[node] = plan[:last]
					return result
				}
			}
		}
	}

	return nil
}
