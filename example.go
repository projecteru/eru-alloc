package main

import (
	"fmt"
	"time"

	"./core"
	"./lock"
	"github.com/coreos/etcd/client"
)

func do(i int, mu *lock.Mutex, api client.KeysAPI, cpu float64, share int) {
	fmt.Println(i, "start")
	mu.Lock()
	fmt.Println(i, "get")
	cpuInfo := core.GetCpuInfo(api, "/h2/cpu")
	host := core.NewHost(cpuInfo, 10)
	cores := host.GetContainerCores(cpu, share)
	fmt.Println(cpu, share, "-----", len(cores), cores)
	fmt.Println(i, "done")
	mu.Unlock()
}

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	cli, _ := client.New(cfg)
	api := client.NewKeysAPI(cli)
	//lock.SetDebugLogger(os.Stdout)
	l := lock.NewMutex(cli, "/test", 10)
	go do(0, l, api, 2.7, -1)
	go do(1, l, api, 1.3, 2)
	go do(2, l, api, 2.5, 2)
	go do(3, l, api, 0.2, 1)
	go do(4, l, api, 3, -1)
	go do(5, l, api, 1.1, -1)
	go do(6, l, api, 1.2, -1)
	for {
		time.Sleep(1 * time.Second)
	}
}
