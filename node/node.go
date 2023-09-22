package node

import (
	"math/rand"
)

// Message Create an enum for messages
type Message int

const (
	Honest      Message = 1
	Default     Message = 0
	Adversarial Message = -1
)

// Node Struct for a single node
type Node struct {
	nodeId         int
	peers          []int
	initialMessage Message
}

// Update the state of the node
func (node *Node) Update() {
	if node.initialMessage == Default {
		// TODO: Implement max counter algorithm
	}
}

// AddPeers to the node
func (node *Node) AddPeers(peers []int) {
	node.peers = peers
}

// Broadcast the message from the node
func (node *Node) Broadcast() Message {
	if node.initialMessage != Default {
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
	var nodes = make(map[int]Node)

	for i := 0; i < totalNodes; i++ {
		if i < numHonestSample {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: Honest}
		} else if i < numHonestSample+numAdversarialSample {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: Adversarial}
		} else {
			nodes[i] = Node{nodeId: i, peers: nil, initialMessage: Default}
		}
	}
	return nodes
}

func generatePeerList(totalNodes int, nodeId int, numPeers int) []int {
	// Get a list of all nodeIds
	nodeIds := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		nodeIds[i] = i
	}

	// Create a list of all possible peers excluding current nodeId
	peers := append(nodeIds[:nodeId], nodeIds[nodeId+1:]...)

	// Randomly sample numPeers nodes from the peer list of a single node
	idx := rand.Perm(len(peers))
	peerList := make([]int, numPeers)
	for i := 0; i < numPeers; i++ {
		peerList[i] = peers[idx[i]]
	}
	return peerList
}

// ConnectNodesToRandomPeers Connect nodes to a defined number of peers
func ConnectNodesToRandomPeers(nodes map[int]Node, numPeers int) map[int]Node {
	totalNodes := len(nodes)

	for i := 0; i < len(nodes); i++ {
		// Get a list of all the nodes
		peerList := generatePeerList(totalNodes, i, numPeers)

		// Update the peer list of a node
		currentNode := nodes[i]
		currentNode.AddPeers(peerList)
		nodes[i] = currentNode
	}
	return nodes
}
