package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	distLock "github.com/mengjinglei/distributed-lock"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	lock := distLock.NewDistLock()
	go lock.Run(*id, *kvport, *cluster, *join)
	time.Sleep(time.Second * 5)
	key := "my-lock"
	for {
		r := rand.Intn(10) + 1
		ticker := time.NewTicker(time.Second * time.Duration(r))
		fmt.Println("=================")
		select {
		case now := <-ticker.C:
			if r%1 == 0 {
				err := lock.LockWithTTL(key, r)
				if err != nil {
					fmt.Printf(">>>>> lock fail  %s, error: %s, time: %d, ttl: %d", key, err.Error(), now.Unix(), r)
					continue
				}
				fmt.Printf(">>>>> lock success  %s, time: %d, ttl: %d", key, now.Unix(), r)
			} else {
				err := lock.Unlock(key)
				if err != nil {
					fmt.Printf(">>>>> unlock fail  %s, error: %s, time: %d, ttl: %d", key, err.Error(), now.Unix(), r)
					continue
				}
				fmt.Printf(">>>>> unlock success  %s, time: %d, ttl: %d", key, now.Unix(), r)
			}

		}
		ticker.Stop()
	}
}
