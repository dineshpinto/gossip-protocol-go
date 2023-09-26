package stf

import (
	"github.com/dineshpinto/gossip-protocol-go/node"
	"testing"
)

func TestEvolveState(t *testing.T) {
	expectedState := 1
	nodes := node.CreateNodes(10, 5, 50)
	state := EvolveState(nodes, 50, 6)
	gotState := state[len(state)-1]
	if gotState != expectedState {
		t.Errorf("Incorrect evolution of state expected (%d), got (%d)",
			expectedState, gotState)
	}
}
