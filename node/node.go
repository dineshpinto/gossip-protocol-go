package node

import (
	"math/rand"
	"sort"
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
	NodeId         int
	Peers          []int
	InitialMessage Message
	MessageCounter map[Message]int
}

func (n *Node) InitializeCounter() {
	n.MessageCounter = map[Message]int{
		Honest:      0,
		Adversarial: 0,
		Default:     0,
	}
}

// Update the state of the node
func (n *Node) Update(messages []Message) {
	if n.InitialMessage == Default {
		for i := 0; i < len(messages); i++ {
			n.MessageCounter[messages[i]] = n.MessageCounter[messages[i]] + 1
		}
	}
}

// AddPeers to the node
func (n *Node) AddPeers(peers []int) {
	n.Peers = peers
}

// Broadcast the message from the node
func (n *Node) Broadcast() Message {
	if n.InitialMessage != Default {
		// Return the initial message for sample nodes
		return n.InitialMessage
	} else {
		// Return the most common message for non-sample nodes
		keys := make([]Message, 0, len(n.MessageCounter))
		for k := range n.MessageCounter {
			keys = append(keys, k)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return n.MessageCounter[keys[i]] < n.MessageCounter[keys[j]]
		})
		return keys[len(keys)-1]
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
			nodes[i] = Node{NodeId: i, Peers: nil, InitialMessage: Honest, MessageCounter: nil}
		} else if i < numHonestSample+numAdversarialSample {
			nodes[i] = Node{NodeId: i, Peers: nil, InitialMessage: Adversarial, MessageCounter: nil}
		} else {
			nodes[i] = Node{NodeId: i, Peers: nil, InitialMessage: Default, MessageCounter: nil}
		}

		// Initialize message counter
		n := nodes[i]
		n.InitializeCounter()
		nodes[i] = n
	}
	return nodes
}

func generatePeerList(totalNodes int, nodeId int, numPeers int) []int {
	// Get a list of all nodeIds
	nodeIds := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		nodeIds[i] = i
	}

	// Create a list of all possible Peers excluding current NodeId
	peers := append(nodeIds[:nodeId], nodeIds[nodeId+1:]...)

	// Randomly sample numPeers nodes from the peer list of a single node
	idx := rand.Perm(len(peers))
	peerList := make([]int, numPeers)
	for i := 0; i < numPeers; i++ {
		peerList[i] = peers[idx[i]]
	}
	return peerList
}

// ConnectNodesToRandomPeers Connect nodes to a defined number of Peers
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
