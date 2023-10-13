package main

import (
	"github.com/dineshpinto/gossip-protocol-go/node"
	"github.com/dineshpinto/gossip-protocol-go/stf"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nodes, _ := node.CreateNodes(6, 4, 1000, false)
		_, _ = stf.EvolveState(nodes, 200, 6, false)
	}
}
