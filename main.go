package main

import (
	"github.com/dineshpinto/gossip-protocol-go/node"
	"github.com/dineshpinto/gossip-protocol-go/stf"
)

func main() {
	// Define network parameters
	numHonestSample := 20
	numAdversarialSample := 15
	numNonSample := 50
	numPeers := 10
	cycles := 50
	// Create network and evolve state
	nodes := node.CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	_ = stf.EvolveState(nodes, cycles, numPeers)
}
