package node

import (
	"math/rand"
	"sort"
)

// Message Create an enum for messages
type Message int

const (
	MessageHonest      Message = 1
	MessageDefault     Message = 0
	MessageAdversarial Message = -1
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
		MessageHonest:      0,
		MessageAdversarial: 0,
		MessageDefault:     0,
	}
}

// Update the state of the node
func (n *Node) Update(messages []Message) {
	if n.InitialMessage == MessageDefault {
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

// CreateNodes Create a set of nodes
func CreateNodes(
	numHonestSample int,
	numAdversarialSample int,
	numNonSample int,
) map[int]Node {
	totalNodes := numHonestSample + numAdversarialSample + numNonSample
	var nodes = make(map[int]Node)

	for nodeId := 0; nodeId < totalNodes; nodeId++ {
		if nodeId < numHonestSample {
			nodes[nodeId] = Node{
				NodeId:         nodeId,
				Peers:          nil,
				InitialMessage: MessageHonest,
				MessageCounter: nil,
			}
		} else if nodeId < numHonestSample+numAdversarialSample {
			nodes[nodeId] = Node{
				NodeId:         nodeId,
				Peers:          nil,
				InitialMessage: MessageAdversarial,
				MessageCounter: nil,
			}
		} else {
			nodes[nodeId] = Node{
				NodeId:         nodeId,
				Peers:          nil,
				InitialMessage: MessageDefault,
				MessageCounter: nil,
			}
		}

		// Initialize message counter
		n := nodes[nodeId]
		n.InitializeCounter()
		nodes[nodeId] = n
	}
	return nodes
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
