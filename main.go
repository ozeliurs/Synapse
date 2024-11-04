package main

import (
	"fmt"
	synapse "synapse/src"
	"time"
)

func createNetwork(name string, nodes []*synapse.Node) *synapse.Network {
	network := synapse.NewNetwork(name)
	for _, node := range nodes {
		node.JoinNetwork(network)
	}
	return network
}

func simulateDataReplication(nodes []*synapse.Node) {
	// Simulate data replication across nodes
	fmt.Println("\n=== Starting Data Replication Test ===")
	testData := map[string]interface{}{
		"user:1":     map[string]string{"name": "Alice", "role": "admin"},
		"user:2":     map[string]string{"name": "Bob", "role": "user"},
		"settings:1": map[string]interface{}{"theme": "dark", "notifications": true},
	}

	// Distribute data across nodes
	for key, value := range testData {
		sourceNode := nodes[len(nodes)%3] // Round-robin distribution
		fmt.Printf("Storing %s through node %s:%d\n", key, sourceNode.IP, sourceNode.Port)
		sourceNode.OPE("PUT", key, value)
		time.Sleep(100 * time.Millisecond)
	}
}

func simulateNetworkPartition(network *synapse.Network, partitionedNode *synapse.Node) {
	fmt.Println("\n=== Simulating Network Partition ===")
	network.RemoveNode(partitionedNode)
	fmt.Printf("Node %s:%d has been partitioned from the network\n", partitionedNode.IP, partitionedNode.Port)
}

func simulateNetworkRecovery(network *synapse.Network, recoveredNode *synapse.Node) {
	fmt.Println("\n=== Simulating Network Recovery ===")
	network.AddNode(recoveredNode)
	fmt.Printf("Node %s:%d has rejoined the network\n", recoveredNode.IP, recoveredNode.Port)
}

func simulateQueryOperations(nodes []*synapse.Node) {
	fmt.Println("\n=== Starting Query Operations ===")
	queries := []struct {
		operation string
		key       string
		value     interface{}
	}{
		{"GET", "user:1", nil},
		{"GET", "user:2", nil},
		{"GET", "settings:1", nil},
		{"PUT", "user:3", map[string]string{"name": "Charlie", "role": "user"}},
		{"GET", "user:3", nil},
	}

	for _, query := range queries {
		sourceNode := nodes[len(nodes)%4] // Round-robin querying
		fmt.Printf("Executing %s operation for key %s through node %s:%d\n",
			query.operation, query.key, sourceNode.IP, sourceNode.Port)
		sourceNode.OPE(query.operation, query.key, query.value)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	// Create multiple nodes with different ports
	nodes := []*synapse.Node{
		synapse.NewNode("127.0.0.1", 12341),
		synapse.NewNode("127.0.0.1", 12342),
		synapse.NewNode("127.0.0.1", 12343),
		synapse.NewNode("127.0.0.1", 12344),
		synapse.NewNode("127.0.0.1", 12345),
	}

	// Start all nodes
	fmt.Println("=== Starting Nodes ===")
	for _, node := range nodes {
		go func(n *synapse.Node) {
			if err := n.Start(); err != nil {
				fmt.Printf("Error starting node %s:%d: %v\n", n.IP, n.Port, err)
			} else {
				fmt.Printf("Node started at %s:%d\n", n.IP, n.Port)
			}
		}(node)
	}

	// Wait for nodes to initialize
	time.Sleep(time.Second)

	// Create two separate networks
	network1 := createNetwork("net1", nodes[:3]) // First 3 nodes
	network2 := createNetwork("net2", nodes[3:]) // Last 2 nodes
	fmt.Printf("\nCreated network1 with %d nodes\n", len(network1.GetNodes()))
	fmt.Printf("Created network2 with %d nodes\n", len(network2.GetNodes()))

	// Simulate initial data replication
	simulateDataReplication(nodes)

	// Let the network stabilize
	time.Sleep(time.Second)

	// Simulate network partition
	simulateNetworkPartition(network1, nodes[0])

	// Simulate operations during partition
	simulateQueryOperations(nodes[1:]) // Exclude partitioned node

	// Simulate network recovery
	time.Sleep(time.Second)
	simulateNetworkRecovery(network1, nodes[0])

	// Final round of queries after recovery
	simulateQueryOperations(nodes)

	// Create cross-network connection
	fmt.Println("\n=== Creating Cross-Network Connection ===")
	nodes[2].JoinNetwork(network2) // Bridge node connecting both networks
	fmt.Printf("Node %s:%d now connects both networks\n", nodes[2].IP, nodes[2].Port)

	// Final test operations across networks
	fmt.Println("\n=== Testing Cross-Network Operations ===")
	testMsg := synapse.NewMessage("FIND", "cross_network_test", "test_value", nodes[0].IP, nodes[4].IP)
	if err := nodes[0].SendMessage(testMsg, nodes[4].GetAddress()); err != nil {
		fmt.Printf("Cross-network message failed: %v\n", err)
	} else {
		fmt.Println("Cross-network message sent successfully")
	}

	// Keep program running to see results
	time.Sleep(5 * time.Second)
	fmt.Println("\n=== Test Scenario Completed ===")
}
