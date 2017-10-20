package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/raft"
	distLock "github.com/mengjinglei/distributed-lock"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	lock, err := distLock.NewDistLock(*id, *kvport, *cluster, *join)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ticker := time.NewTicker(time.Second * rand.Intn(10))
		case now := <-ticker.C:
			if rc.node.Status().SoftState.RaftState == raft.StateLeader {
				fmt.Printf("%s start to check ttl\n", now.String())
				rc.checkTTL(now)
			}
		}
		ticker.Stop()
}
