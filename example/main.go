package main

import (
	"flag"
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"

	distLock "github.com/mengjinglei/distributed-lock"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	lock := distLock.NewDistLock(*id, *kvport, *cluster, *join)
	time.Sleep(time.Second * 5)
	key := "my-lock"
	for {
		r := rand.Intn(10) + 1
		ticker := time.NewTicker(time.Second * time.Duration(r))
		select {
		case now := <-ticker.C:
			if r%1 == 0 {
				err := lock.LockWithTTL(key, r)
				if err != nil {
					log.Printf("lock fail  %s, error: %s, time: %d, ttl: %d, id: %d\n", key, err.Error(), now.Unix(), r, id)
					continue
				}
				log.Printf("lock success  %s, time: %d, ttl: %d, id: %d\n", key, now.Unix(), r, id)
			} else {
				err := lock.Unlock(key)
				if err != nil {
					log.Printf("unlock fail  %s, error: %s, time: %d, ttl: %d, id: %d\n", key, err.Error(), now.Unix(), r, id)
					continue
				}
				log.Printf("unlock success  %s, time: %d, ttl: %d, id: %d\n", key, now.Unix(), r, id)
			}

		}
		ticker.Stop()
	}
}
