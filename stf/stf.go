package stf

import (
	"gossip_protocol_go/node"
	"log"
)

// EvolveState Evolve the state of the system in time
func EvolveState(nodes map[int]node.Node, cycles int, numPeers int) []int {
	messageQueue := make(map[int][]node.Message)
	var nonSampleBroadcasts []int

	for cycle := 0; cycle < cycles; cycle++ {
		var _nonSampleBroadcasts []int
		nodes = node.ConnectNodesToRandomPeers(nodes, numPeers)
		for i := 0; i < len(nodes); i++ {
			n := nodes[i]
			// Extract the broadcast message from each node
			msg := n.Broadcast()
			// Save the message for further analysis
			if n.InitialMessage == node.MessageDefault {
				_nonSampleBroadcasts = append(_nonSampleBroadcasts, int(msg))
			}
			// Don't add the message to the queue if the node has nothing to say
			if msg == node.MessageDefault {
				continue
			}
			// Add the node message to a queue with its peers
			for _, nodeId := range n.Peers {
				messageQueue[nodeId] = append(messageQueue[nodeId], msg)
			}
			nodes[i] = n
		}
		// log.Println("message queue", messageQueue)
		// Update all nodes with the messages in the queue
		for nodeId, messages := range messageQueue {
			n := nodes[nodeId]
			n.Update(messages)
			nodes[nodeId] = n
		}
		// Clear message queue
		messageQueue = make(map[int][]node.Message)
		log.Println("average value", average(_nonSampleBroadcasts))
		nonSampleBroadcasts = append(nonSampleBroadcasts, _nonSampleBroadcasts...)
	}
	return nonSampleBroadcasts
}

// Calculate the average value of an integer array
func average(arr []int) float64 {
	arrSum := 0
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		arrSum += arr[i]
	}
	return float64(arrSum) / float64(arrLen)
}
