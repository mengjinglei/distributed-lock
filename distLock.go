// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package distlock

import (
	"strings"

	"github.com/coreos/etcd/raft/raftpb"
)

type DistLock struct {
	rc          *raftNode
	proposeC    chan string
	confChangeC chan raftpb.ConfChange

	id      int
	cluster string
	port    int
	join    bool
}

func (dl *DistLock) Lock(key string) error {

	return dl.rc.Lock(key)
}

func (dl *DistLock) LockWithTTL(key string, ttl int) error {

	return dl.rc.LockWithTTL(key, ttl)
}

func (dl *DistLock) Unlock(key string) error {

	return dl.rc.Unlock(key)
}

func (dl *DistLock) IsLeader() bool {
	return dl.rc.IsLeader()
}

func (sl *DistLock) Stop() {
	close(sl.proposeC)
	close(sl.confChangeC)
}

func NewDistLock(id, port int, cluster string, join bool) *DistLock {

	dl := &DistLock{
		proposeC:    make(chan string),
		confChangeC: make(chan raftpb.ConfChange),
	}

	go func() {
		rc := newRaftNode(id, strings.Split(cluster, ","), join, dl.proposeC, dl.confChangeC)
		dl.rc = rc

	}()

	return dl
}
