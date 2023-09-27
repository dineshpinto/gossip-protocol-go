package main

import (
	"github.com/dineshpinto/gossip-protocol-go/node"
	"github.com/dineshpinto/gossip-protocol-go/stf"
	"log"
)

func main() {
	// Define network parameters
	numHonestSample := 20
	numAdversarialSample := 15
	numNonSample := 2000
	numPeers := 5
	cycles := 50
	// Create network and evolve state
	nodes, err := node.CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stf.EvolveState(nodes, cycles, numPeers)
	if err != nil {
		log.Fatal(err)
	}
}
