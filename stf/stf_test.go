package stf

import (
	"github.com/dineshpinto/gossip-protocol-go/node"
	"testing"
)

func TestEvolveState(t *testing.T) {
	wantState := 1
	nodes, err := node.CreateNodes(10, 1, 50)
	if err != nil {
		t.Errorf("Error creating nodes %s", err)
	}
	state, err := EvolveState(nodes, 50, 6)
	if err != nil {
		t.Errorf("Error evolving state %s", err)
	}
	gotState := state[len(state)-1]
	if gotState != wantState {
		t.Errorf("Incorrect evolution of state expected (%d), got (%d)",
			wantState, gotState)
	}
}
