package synapse_test

import (
	synapse "synapse/src"
	"testing"
)

func TestNewNode(t *testing.T) {
    node := synapse.NewNode("127.0.0.1", 8080)

    if node.IP != "127.0.0.1" {
        t.Errorf("Expected IP '127.0.0.1', got %s", node.IP)
    }
    if node.Port != 8080 {
        t.Errorf("Expected Port 8080, got %d", node.Port)
    }
}

func TestNodeGetAddress(t *testing.T) {
    node := synapse.NewNode("127.0.0.1", 8080)
    expected := "127.0.0.1:8080"

    if addr := node.GetAddress(); addr != expected {
        t.Errorf("Expected address %s, got %s", expected, addr)
    }
}

func TestNodeJoinNetwork(t *testing.T) {
    node := synapse.NewNode("127.0.0.1", 8080)
    network := synapse.NewNetwork("test-network")

    node.JoinNetwork(network)

    if len(node.Networks) != 1 {
        t.Errorf("Expected 1 network, got %d", len(node.Networks))
    }
    if len(network.GetNodes()) != 1 {
        t.Errorf("Expected 1 node in network, got %d", len(network.GetNodes()))
    }
}
