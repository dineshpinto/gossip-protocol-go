package node

import (
	"testing"
)

func TestCreateNodes(t *testing.T) {
	wantHonestSample := 10
	wantAdversarialSample := 5
	wantNonSample := 50
	nodes, err := CreateNodes(wantHonestSample, wantAdversarialSample, wantNonSample, true)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}

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
	wantPeers := 7
	nodes, err := CreateNodes(10, 5, 50, true)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}
	nodes, err = ConnectNodesToRandomPeers(nodes, wantPeers)
	if err != nil {
		t.Errorf("Error connecting nodes %s", err)
	}
	gotPeers := len(nodes[0].Peers)
	if gotPeers != wantPeers {
		t.Errorf("Incorrect connection of peers expected (%d) got (%d)",
			wantPeers, gotPeers)
	}
}

func TestSampleNode_Broadcast(t *testing.T) {
	nodes, err := CreateNodes(10, 5, 50, true)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}
	node := nodes[0]
	wantBroadcast := node.InitialMessage
	gotBroadcast := node.Broadcast()
	if gotBroadcast != wantBroadcast {
		t.Errorf("Incorrect message broadcast by sample Node expected (%d) "+
			"got (%d)",
			wantBroadcast, gotBroadcast)
	}
}

func TestNonSampleNode_Broadcast(t *testing.T) {
	nodes, err := CreateNodes(10, 5, 50, true)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}
	node := nodes[len(nodes)-1]
	wantBroadcast := node.InitialMessage
	gotBroadcast := node.Broadcast()
	if gotBroadcast != wantBroadcast {
		t.Errorf("Incorrect message broadcast by non-sample Node expected "+
			"(%d) got (%d)",
			wantBroadcast, gotBroadcast)
	}
}

func TestNode_Update(t *testing.T) {
	nodes, err := CreateNodes(10, 5, 50, true)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}
	node := nodes[len(nodes)-1]
	wantMessageCount := node.MessageCounter[MessageHonest] + 1
	node.Update([]Message{MessageHonest})
	gotMessageCount := node.MessageCounter[MessageHonest]
	if gotMessageCount != wantMessageCount {
		t.Errorf("Incorrect message update by node expected (%d) got (%d)",
			wantMessageCount, gotMessageCount)
	}
}

func BenchmarkCreateNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = CreateNodes(6, 4, 1000, false)
	}
}

func BenchmarkConnectNodesToRandomPeers(b *testing.B) {
	nodes, _ := CreateNodes(6, 4, 1000, false)
	for i := 0; i < b.N; i++ {
		_, _ = ConnectNodesToRandomPeers(nodes, 7)
	}
}
