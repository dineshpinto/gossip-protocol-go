package main

func main() {
	// Define network parameters
	numHonestSample := 20
	numAdversarialSample := 15
	numNonSample := 50
	numPeers := 10
	cycles := 50
	// Create network and evolve state
	nodes := CreateNodes(numHonestSample, numAdversarialSample, numNonSample)
	_ = EvolveState(nodes, cycles, numPeers)
}
