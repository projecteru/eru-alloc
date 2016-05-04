package core

type Host struct {
	full     map[string]int
	fragment map[string]int
	share    int
}

func NewHost(cpuInfo map[string]int, share int) *Host {
	host := &Host{
		share:    share,
		full:     map[string]int{},
		fragment: map[string]int{},
	}
	for no, pieces := range cpuInfo {
		if pieces == share {
			host.full[no] = pieces
		} else {
			host.fragment[no] = pieces
		}
	}

	return host
}

func (self *Host) min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (self *Host) calcuatePiecesCores(full int, fragment int, maxShareCore int) {
	if maxShareCore == -1 {
		maxShareCore = len(self.full) - len(self.fragment) - full // 减枝，M == N 的情况下预留至少一个 full 量的核数
	}

	fullResultNum := len(self.full) / full
	fragmentPiecesTotal := 0
	for _, pieces := range self.fragment {
		fragmentPiecesTotal += pieces
	}
	fragmentResultNum := fragmentPiecesTotal / fragment
	baseLine := self.min(fullResultNum, fragmentResultNum)

	num := 0
	for i := 1; i < maxShareCore+1; i++ {
		if len(self.fragment) > i {
			continue
		}
		fullResultNum = (len(self.full) - i) / full
		fragmentResultNum = 0
		for j := 0; j < i; j++ {
			fragmentResultNum += self.share / fragment
		}
		canDeployNum := self.min(fullResultNum, fragmentResultNum)
		if canDeployNum > baseLine {
			num = i
			baseLine = canDeployNum
		}
	}

	count := 0
	for no, pieces := range self.full {
		if count == num {
			break
		}
		self.fragment[no] = pieces
		count += 1
		delete(self.full, no)
	}
}

func (self *Host) GetContainerCores(num float64, maxShareCore int) []map[string]int {
	num = num * float64(self.share)
	full := int(num) / self.share
	fragment := int(num) % self.share
	result := []map[string]int{}

	if full == 0 {
		if maxShareCore == -1 {
			maxShareCore = len(self.full)
		}
		for no, pieces := range self.full {
			if len(self.fragment) == maxShareCore {
				break
			}
			self.fragment[no] = pieces
			delete(self.full, no)
		}
		fragmentResult := self.getFragmentResult(fragment)
		for _, no := range fragmentResult {
			result = append(result, map[string]int{no: fragment})
		}
		return result
	}

	if fragment == 0 {
		n := len(self.full) / full
		for i := 0; i < n; i++ {
			fullResult := self.getFullResult(full)
			result = append(result, fullResult)
		}
		return result
	}

	// 算出最优的碎片核和整数核组合
	self.calcuatePiecesCores(full, fragment, maxShareCore)
	fragmentResult := self.getFragmentResult(fragment)
	for _, no := range fragmentResult {
		fullResult := self.getFullResult(full)
		if len(fullResult) != full { // 可能整数核不够用了结果并不一定可靠必须再判断一次
			return result // 减枝这时候整数核一定不够用了，直接退出，这样碎片核和整数核的计算就完成了
		}
		fullResult[no] = fragment
		result = append(result, fullResult)
	}
	return result
}

func (self *Host) getFragmentResult(fragment int) []string {
	result := []string{}
	for no, pieces := range self.fragment {
		for i := 0; i < pieces/fragment; i++ {
			result = append(result, no)
		}
	}
	return result
}

func (self *Host) getFullResult(full int) map[string]int {
	result := map[string]int{}
	for no, pieces := range self.full {
		result[no] = pieces   // 分配一整个核
		delete(self.full, no) // 干掉这个可用资源
		if len(result) == full {
			break
		}
	}
	return result
}