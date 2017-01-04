//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"math/rand"
	"time"

	"github.com/gogo/protobuf/proto"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func SetSeed(seed int64) {
	r = rand.New(rand.NewSource(seed))
}

func RandomPerson() proto.Message {
	return NewPopulatedPerson(r, true)
}

func RandomSrcTree() proto.Message {
	pops := []*SrcTree{IoUtilSrcTree, PathSrcTree, RuntimeSrcTree, SyscallSrcTree}
	return pops[r.Intn(4)]
}

func RandomTypewriterPrison() proto.Message {
	return NewPopulatedTypewriterPrison(r, true)
}

func RandomPuddingMilkshake() proto.Message {
	return NewPopulatedPuddingMilkshake(r, true)
}
