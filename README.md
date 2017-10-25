# distributed-lock

distributed-lock is an distributed lock written in go by using etcd's [raft library](github.com/coreos/etcd/raft). It provides simple way to use distributed lock without an external etcd cluster.  

## Getting Started

### Install

```
go get github.com/mengjinglei/distributed-lock
```

### Usage

```

import (
    distLock "github.com/mengjinglei/distributed-lock"
)

// Create a distributed lock
distLock,err := distLock.NewDistLock(id, port int, peers string, join bool)
if err != nil{
    fmt.Println("Create distributed lock fail: %s",err.Error())
    return
}

// acquire a lock with TTL(int, number of Seconds)
err := distLock.LockWithTTL(LockKey, TTL)
if err != nil{
    fmt.Println("acquire lock %s fail: %s",LockKey, err.Error())
    return
}

// acquire a lock with TTL = 60 seconds
err := distLock.Lock(LockKey)
if err != nil{
    fmt.Println("acquire lock %s fail: %s",LockKey, err.Error())
    return
}

// release a lock with LockKey
err := distLock.Unlock(LockKey)
if err != nil{
    fmt.Println("acquire lock %s fail: %s",LockKey, err.Error())
    return
}

// check if node is leader
isLeader := distLock.IsLeader()

// stop a lock
distLock.Stop()

```
