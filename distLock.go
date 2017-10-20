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
	rc *raftNode
}

func (dl *DistLock) Lock(key string) error {

	return dl.rc.Lock(key)
}

func (dl *DistLock) LockWithTTL(key string, ttl int) error {

	return dl.rc.LockWithTTL(key, ttl)
}

func NewDistLock(id, port int, cluster string, join bool) {

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	// raft provides a commit stream for the proposals from the http api

	errorC, rc := newRaftNode(id, strings.Split(cluster, ","), join, proposeC, confChangeC)

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(rc.kvStore, port, confChangeC, errorC)
}
