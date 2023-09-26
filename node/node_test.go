package node

import (
	"testing"
)

func TestCreateNodes(t *testing.T) {
	wantHonestSample := 10
	wantAdversarialSample := 5
	wantNonSample := 50
	nodes := CreateNodes(wantHonestSample, wantAdversarialSample, wantNonSample)

	gotHonestSample := 0
	gotAdversarialSample := 0
	gotNonSample := 0
	for _, node := range nodes {
		if node.InitialMessage == MessageHonest {
			gotHonestSample += 1
		} else if node.InitialMessage == MessageAdversarial {
			gotAdversarialSample += 1
		} else if node.InitialMessage == MessageDefault {
			gotNonSample += 1
		}
	}

	if (gotHonestSample != wantHonestSample) ||
		(gotAdversarialSample != wantAdversarialSample) ||
		(gotNonSample != wantNonSample) {
		t.Errorf("Incorrect initialization of nodes expected (%d, %d, %d), "+
			"got (%d, %d, %d)", wantHonestSample, wantAdversarialSample,
			wantNonSample, gotHonestSample, gotAdversarialSample, gotNonSample)
	}
}

func TestConnectNodesToRandomPeers(t *testing.T) {
	expectedPeers := 7
	nodes := CreateNodes(10, 5, 50)
	nodes = ConnectNodesToRandomPeers(nodes, expectedPeers)
	gotPeers := len(nodes[0].Peers)

	if gotPeers != expectedPeers {
		t.Errorf("Incorrect connection of peers expected (%d) got (%d)",
			expectedPeers, gotPeers)
	}
}

func TestNode_Broadcast(t *testing.T) {
	nodes := CreateNodes(10, 5, 50)
	// Test Sample Node
	sampleNode := nodes[0]
	wantSampleBroadcast := MessageHonest
	gotSampleBroadcast := sampleNode.Broadcast()

	if gotSampleBroadcast != wantSampleBroadcast {
		t.Errorf("Incorrect message broadcast by sample Node expected (%d) got (%d)",
			wantSampleBroadcast, gotSampleBroadcast)
	}

	// Test Non Sample Node
	nonSampleNode := nodes[len(nodes)-1]
	wantNonSampleBroadcast := MessageDefault
	gotNonSampleBroadcast := nonSampleNode.Broadcast()

	if gotNonSampleBroadcast != wantNonSampleBroadcast {
		t.Errorf("Incorrect message broadcast by non-sample Node expected "+
			"(%d) got (%d)",
			wantNonSampleBroadcast, gotNonSampleBroadcast)
	}
}

func TestNode_Update(t *testing.T) {
	nodes := CreateNodes(10, 5, 50)
	node := nodes[len(nodes)-1]

	wantMessageCount := node.MessageCounter[MessageHonest] + 1
	node.Update([]Message{MessageHonest})
	gotMessageCount := node.MessageCounter[MessageHonest]

	if gotMessageCount != wantMessageCount {
		t.Errorf("Incorrect message update by node expected (%d) got (%d)",
			wantMessageCount, gotMessageCount)
	}
}
