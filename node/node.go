package node

import (
	"math/rand"
)

// Message Create an enum for messages
type Message int

const (
	HONEST      Message = 1
	DEFAULT     Message = 0
	ADVERSARIAL Message = -1
)

// Node Struct for a single node
type Node struct {
	nodeId         int
	peers          []int
	initialMessage Message
}

// Update the state of the node
func (node Node) Update() {
	if node.initialMessage == DEFAULT {
		// TODO: Implement max counter algorithm
	}
}

// AddPeers to the node
func (node Node) AddPeers(peers []int) {
	node.peers = peers
}

// Broadcast the message from the node
func (node Node) Broadcast() Message {
	if node.initialMessage != DEFAULT {
		return node.initialMessage
	} else {
		// TODO: Return most common element
		return node.initialMessage
	}
}

// CreateNodes Create a set of nodes
func CreateNodes(
	numHonestSample int,
	numAdversarialSample int,
	numNonSample int,
) map[int]Node {

	totalNodes := numHonestSample + numAdversarialSample + numNonSample
	var nodes map[int]Node
	nodes = make(map[int]Node)

	for i := 0; i < totalNodes; i++ {
		if i < numHonestSample {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: HONEST}
		} else if i < numHonestSample+numAdversarialSample {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: ADVERSARIAL}
		} else {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: DEFAULT}
		}
	}
	return nodes
}

// ConnectNodesToRandomPeers Connect nodes to a defined number of peers
func ConnectNodesToRandomPeers(nodes map[int]Node, numPeers int) map[int]Node {
	totalNodes := len(nodes)

	for nodeId, node := range nodes {
		// Get a list of all the nodes
		// TODO: This should sit outside, but the sampling breaks if it does
		nodeIds := make([]int, totalNodes)
		for i := 0; i < totalNodes; i++ {
			nodeIds[i] = i
		}

		// Randomly sample numPeers nodes from the peer list of a single node
		peers := append(nodeIds[:nodeId], nodeIds[nodeId+1:]...)
		idx := rand.Perm(len(peers))
		peerSample := make([]int, numPeers)
		for i := 0; i < numPeers; i++ {
			peerSample[i] = peers[idx[i]]
		}
		// TODO: This does not work correctly
		node.AddPeers(peerSample)
	}
	return nodes
}
