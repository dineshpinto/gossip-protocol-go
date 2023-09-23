package stf

import (
	"fmt"
	"gossip_protocol_go/node"
)

func EvolveState(nodes map[int]node.Node, cycles int, numPeers int) []int {
	messageQueue := make(map[int][]node.Message)
	var nonSampleBroadcasts []int

	for cycle := 0; cycle < cycles; cycle++ {
		var _nonSampleBroadcasts []int

		nodes = node.ConnectNodesToRandomPeers(nodes, numPeers)

		for i := 0; i < len(nodes); i++ {
			n := nodes[i]
			msg := n.Broadcast()

			if n.InitialMessage == node.MessageDefault {
				_nonSampleBroadcasts = append(_nonSampleBroadcasts, int(msg))
			}
			if msg == node.MessageDefault {
				continue
			}
			for nodeId := range n.Peers {
				messageQueue[nodeId] = append(messageQueue[nodeId], msg)
			}
			nodes[i] = n
		}

		for nodeId, messages := range messageQueue {
			n := nodes[nodeId]
			n.Update(messages)
			nodes[nodeId] = n
		}
		// Clear message queue
		messageQueue = make(map[int][]node.Message)
		fmt.Println(average(_nonSampleBroadcasts))
		nonSampleBroadcasts = append(nonSampleBroadcasts, _nonSampleBroadcasts...)
	}
	return nonSampleBroadcasts
}

func average(arr []int) float64 {
	arrSum := 0
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		arrSum += arr[i]
	}
	return float64(arrSum) / float64(arrLen)
}
