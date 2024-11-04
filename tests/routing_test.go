package synapse_test

import (
	synapse "synapse/src"
	"testing"
)

func TestRoutingTableOperations(t *testing.T) {
    rt := synapse.NewRoutingTable()

    // Test UpdateRoute and GetNextHop
    rt.UpdateRoute("testKey", "node1:8080")

    if hop := rt.GetNextHop("testKey"); hop != "node1:8080" {
        t.Errorf("Expected next hop 'node1:8080', got %s", hop)
    }

    // Test IsResponsible
    if !rt.IsResponsible("testKey") {
        t.Error("Expected to be responsible for 'testKey'")
    }

    if rt.IsResponsible("nonexistentKey") {
        t.Error("Expected not to be responsible for 'nonexistentKey'")
    }
}
