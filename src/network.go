package synapse

import "sync"

type Network struct {
	ID    string
	Nodes []*Node
	mu    sync.RWMutex
}

func NewNetwork(id string) *Network {
	return &Network{
		ID:    id,
		Nodes: make([]*Node, 0),
	}
}

func (n *Network) AddNode(node *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Nodes = append(n.Nodes, node)
}

func (net *Network) RemoveNode(node *Node) {
	net.mu.Lock()
	defer net.mu.Unlock()

	for i, n := range net.Nodes {
		if n == node {
			net.Nodes = append(net.Nodes[:i], net.Nodes[i+1:]...)
			return
		}
	}
}

func (n *Network) GetNodes() []*Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.Nodes
}
