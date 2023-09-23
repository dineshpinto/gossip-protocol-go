package main

import (
	"gossip_protocol_go/node"
	"gossip_protocol_go/stf"
	"log"
)

func main() {
	// Define network parameters
	numHonestSample := 3
	numAdversarialSample := 1
	numNonSample := 6
	numPeers := 3
	cycles := 50
	// Create network and evolve state
	var states []int
	nodes := node.CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	states = stf.EvolveState(nodes, cycles, numPeers)
	//fmt.Println(nodes)
	log.Println(states)
}
