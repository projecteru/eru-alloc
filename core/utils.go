package core

import (
	"encoding/json"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func GetCpuInfo(api client.KeysAPI, path string) map[string]int {
	resp, err := api.Get(context.Background(), path, &client.GetOptions{Recursive: true})
	if err != nil {
		// bla bla
		return nil
	}

	var cpuInfo map[string]int
	if err := json.Unmarshal([]byte(resp.Node.Value), &cpuInfo); err != nil {
		// bla bla
		return nil
	}
	return cpuInfo
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
