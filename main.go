package main

import (
	"fmt"
	"gossip_protocol_go/node"
	"gossip_protocol_go/stf"
)

func main() {
	numHonestSample := 3
	numAdversarialSample := 1
	numNonSample := 6
	numPeers := 3
	cycles := 50

	var states []int
	nodes := node.CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	states = stf.EvolveState(nodes, cycles, numPeers)

	//fmt.Println(nodes)
	fmt.Println(states)
}
