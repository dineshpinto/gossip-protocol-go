package stf

import "gossip_protocol_go/node"

func EvolveState(nodes map[int]node.Node, cycles int, numPeers int) []int {
	messageQueue := make(map[int][]node.Message)
	var nonSampleBroadcasts []int

	for cycle := 0; cycle < cycles; cycle++ {
		nodes = node.ConnectNodesToRandomPeers(nodes, numPeers)
		var _nonSampleBroadcasts []int

		for i := 0; i < len(nodes); i++ {
			n := nodes[i]
			msg := n.Broadcast()

			if n.InitialMessage == node.Default {
				_nonSampleBroadcasts = append(_nonSampleBroadcasts, int(msg))
			}

			if msg == node.Default {
				continue
			}

			for nodeId := range n.Peers {
				messageQueue[nodeId] = append(messageQueue[nodeId], msg)
			}

			// Clear message queue
			messageQueue = make(map[int][]node.Message)
		}
		nonSampleBroadcasts = append(nonSampleBroadcasts, _nonSampleBroadcasts...)
	}
	return nonSampleBroadcasts
}
