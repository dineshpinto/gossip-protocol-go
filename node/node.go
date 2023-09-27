package node

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

// Message Create an enum for messages
type Message int

const (
	MessageAdversarial Message = iota - 1
	MessageDefault
	MessageHonest
)

// Node Struct for a single node
type Node struct {
	NodeId         int
	Peers          []int
	InitialMessage Message
	MessageCounter map[Message]int
}

// Update the state of the node
func (n *Node) Update(messages []Message) {
	if n.InitialMessage == MessageDefault {
		for i := 0; i < len(messages); i++ {
			n.MessageCounter[messages[i]] = n.MessageCounter[messages[i]] + 1
		}
	}
}

// Broadcast the message from the node
func (n *Node) Broadcast() Message {
	if n.InitialMessage != MessageDefault {
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
		// log.Println(len(n.MessageCounter), keys[len(keys)-1])
		return keys[len(keys)-1]
	}
}

// addPeers to the node
func (n *Node) addPeers(peers []int) {
	n.Peers = peers
}

// Create a newNode with a node ID and initial message
func newNode(nodeId int, initialMessage Message) Node {
	messageDefaultCounts := 0
	messageHonestCounts := 0
	messageAdversarialCounts := 0
	if initialMessage == MessageDefault {
		messageDefaultCounts += 1
	} else if initialMessage == MessageHonest {
		messageHonestCounts += 1
	} else if initialMessage == MessageAdversarial {
		messageAdversarialCounts += 1
	}
	counter := map[Message]int{
		MessageHonest:      messageHonestCounts,
		MessageAdversarial: messageAdversarialCounts,
		MessageDefault:     messageDefaultCounts,
	}
	return Node{
		NodeId:         nodeId,
		Peers:          nil,
		InitialMessage: initialMessage,
		MessageCounter: counter,
	}
}

// CreateNodes Create a set of nodes
func CreateNodes(
	numHonestSample int,
	numAdversarialSample int,
	numNonSample int,
) (map[int]Node, error) {
	totalNodes := numHonestSample + numAdversarialSample + numNonSample
	log.Printf(
		"Setting up network with %d sample nodes (honest = %d, "+
			"adversarial = %d) and %d non sample nodes\n",
		numHonestSample+numAdversarialSample, numHonestSample, numAdversarialSample,
		numNonSample)
	nodes := make(map[int]Node, totalNodes)
	for nodeId := 0; nodeId < totalNodes; nodeId++ {
		if nodeId < numHonestSample {
			nodes[nodeId] = newNode(nodeId, MessageHonest)
		} else if nodeId < numHonestSample+numAdversarialSample {
			nodes[nodeId] = newNode(nodeId, MessageAdversarial)
		} else {
			nodes[nodeId] = newNode(nodeId, MessageDefault)
		}
	}

	// Check node creation
	actualHonest, actualAdversarial, actualDefault := 0, 0, 0
	for i := 0; i < len(nodes); i++ {
		// Check if nodes have at least one message
		if nodes[i].MessageCounter[MessageHonest] != 1 &&
			nodes[i].MessageCounter[MessageDefault] != 1 &&
			nodes[i].MessageCounter[MessageAdversarial] != 1 {
			return nil, fmt.Errorf("error creating nodes, no message")
		}
		// Check if the number of nodes is correct
		if nodes[i].InitialMessage == MessageHonest {
			actualHonest += 1
		} else if nodes[i].InitialMessage == MessageAdversarial {
			actualAdversarial += 1
		} else if nodes[i].InitialMessage == MessageDefault {
			actualDefault += 1
		}

	}
	if len(nodes) != totalNodes ||
		actualHonest != numHonestSample ||
		actualAdversarial != numAdversarialSample ||
		actualDefault != numNonSample {
		return nil, fmt.Errorf("error creating nodes, wrong number of nodes")
	}
	return nodes, nil
}

// Generate a peer list for a given node
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
func ConnectNodesToRandomPeers(nodes map[int]Node, numPeers int) (map[int]Node, error) {
	totalNodes := len(nodes)
	for i := 0; i < len(nodes); i++ {
		// Get a list of all the nodes
		peerList := generatePeerList(totalNodes, i, numPeers)
		// Update the peer list of a node
		currentNode := nodes[i]
		currentNode.addPeers(peerList)
		nodes[i] = currentNode

		if len(nodes[i].Peers) != numPeers {
			return nil, fmt.Errorf("error connecting nodes to peers")
		}
	}
	return nodes, nil
}
