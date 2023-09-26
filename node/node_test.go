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
