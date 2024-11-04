package synapse_test

import (
	synapse "synapse/src"
	"testing"
)

func TestNewNetwork(t *testing.T) {
    network := synapse.NewNetwork("test-network")

    if network.ID != "test-network" {
        t.Errorf("Expected network ID 'test-network', got %s", network.ID)
    }

    if len(network.GetNodes()) != 0 {
        t.Errorf("Expected empty nodes list, got %d nodes", len(network.GetNodes()))
    }
}

func TestAddRemoveNode(t *testing.T) {
    network := synapse.NewNetwork("test-network")
    node := synapse.NewNode("127.0.0.1", 8080)

    network.AddNode(node)

    if len(network.GetNodes()) != 1 {
        t.Errorf("Expected 1 node, got %d nodes", len(network.GetNodes()))
    }

    network.RemoveNode(node)

    if len(network.GetNodes()) != 0 {
        t.Errorf("Expected 0 nodes after removal, got %d nodes", len(network.GetNodes()))
    }
}
