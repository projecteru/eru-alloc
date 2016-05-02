package main

import (
	"fmt"
	"time"

	"./lock"
	"github.com/coreos/etcd/client"
)

func do(i int, mu *lock.Mutex) {
	fmt.Println(i, "start")
	mu.Lock()
	fmt.Println(i, "get")
	time.Sleep(5 * time.Second)
	fmt.Println(i, "done")
	mu.Unlock()
}

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	cli, _ := client.New(cfg)
	//lock.SetDebugLogger(os.Stdout)
	l := lock.NewMutex(cli, "/test", 10)
	for i := 0; i < 5; i++ {
		go do(i, l)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
