package main

import (
	"flag"
	"github.com/dineshpinto/gossip-protocol-go/node"
	"github.com/dineshpinto/gossip-protocol-go/stf"
	"log"
)

func main() {
	// Define network parameters
	numHonestSample := flag.Int(
		"hon", 20, "number of honest sample nodes")
	numAdversarialSample := flag.Int(
		"adv", 15, "number of adversarial sample nodes")
	numNonSample := flag.Int(
		"ns", 1000, "number of non sample nodes")
	numPeers := flag.Int(
		"peers", 5, "number of peers per node")
	cycles := flag.Int(
		"cycles", 50, "number of cycles to evolve the network")
	flag.Parse()
	// Create network nodes
	nodes, err := node.CreateNodes(
		*numHonestSample,
		*numAdversarialSample,
		*numNonSample,
		true,
	)
	if err != nil {
		log.Fatal(err)
	}
	// Evolve network for a given number of cycles
	_, err = stf.EvolveState(nodes, *cycles, *numPeers, true)
	if err != nil {
		log.Fatal(err)
	}
}
