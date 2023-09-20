package main

import (
	"fmt"
	"gossip_protocol_go/node"
)

func main() {
	nodes := node.CreateNodes(3, 2, 5)
	nodes = node.ConnectNodesToRandomPeers(nodes, 2)
	fmt.Println(nodes)
}
