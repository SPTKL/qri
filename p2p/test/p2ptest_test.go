package p2ptest

import (
	"context"
	"testing"
)

// Ensure that when we use ConnectNodes, we are creating a basic connection
// between two nodes
// - we have connections to each peer
// - we have the addrs of each peer
// - we have a tag on each peer
func TestConnectNodes(t *testing.T) {
	ctx := context.Background()
	f := NewTestNodeFactory(NewTestableNode)
	testNodes, err := NewTestDirNetwork(ctx, f)
	if err != nil {
		t.Error(err)
		return
	}

	if err := ConnectNodes(ctx, testNodes); err != nil {
		t.Error(err)
	}

	for _, node := range testNodes {
		// test that each conn has a connection to at least one peer id
		pid := node.SimplePeerInfo().ID
		for _, rnode := range testNodes {
			rpid := rnode.SimplePeerInfo().ID
			// dont need to check for connections to self
			if pid == rpid {
				continue
			}
			conns := node.Host().Network().ConnsToPeer(rpid)
			if len(conns) == 0 {
				t.Errorf("node %s has no connections to node %s", pid, rpid)
			}
			addrs := node.Host().Peerstore().Addrs(rpid)
			if len(addrs) == 0 {
				t.Errorf("node %s has no addrs for node %s", pid, rpid)
			}
			tag := node.Host().ConnManager().GetTagInfo(rpid)
			if tag == nil {
				t.Errorf("node %s has not tag info on node %s", pid, rpid)
			}
		}
	}
}
