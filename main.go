package main

import (
	"fmt"
	"gossip_protocol_go/node"
)

func main() {
	numHonestSample := 3
	numAdversarialSample := 2
	numNonSample := 5
	numPeers := 3

	nodes := node.CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	fmt.Println(nodes)
	nodes = node.ConnectNodesToRandomPeers(nodes, numPeers)
	fmt.Println(nodes)
}
