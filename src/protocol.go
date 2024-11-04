package synapse

import "fmt"

type SynapseProtocol struct {
	MRR int
	TTL int
}

func NewSynapseProtocol() *SynapseProtocol {
	return &SynapseProtocol{
		MRR: 3,
		TTL: 10,
	}
}

func (p *SynapseProtocol) HandleMessage(msg *Message, node *Node) {
	switch msg.Type {
	case "FIND":
		p.handleFind(msg, node)
	case "FOUND":
		p.handleFound(msg, node)
	case "JOIN":
		p.handleJoin(msg, node)
	}
}

func (p *SynapseProtocol) HandleOPE(msg *Message, node *Node) {
	findMsg := NewMessage("FIND", msg.Key, msg.Value, node.IP, "")
	findMsg.TTL = p.TTL
	findMsg.MRR = p.MRR
	findMsg.Tag = msg.Tag

	for _, network := range node.Networks {
		for _, peer := range network.Nodes {
			node.SendMessage(findMsg, peer.GetAddress())
		}
	}
}

func (p *SynapseProtocol) handleFind(msg *Message, node *Node) {
	if msg.TTL <= 0 || node.TagManager.IsProcessed(msg.Tag) {
		return
	}

	node.TagManager.PushTag(msg.Tag)
	msg.TTL--

	// Check if responsible
	if node.Routing.IsResponsible(msg.Key) {
		foundMsg := NewMessage("FOUND", msg.Key, msg.Value, node.IP, msg.Source)
		node.SendMessage(foundMsg, msg.Source)
		return
	}

	// Forward to other networks
	nextHop := node.Routing.GetNextHop(msg.Key)
	if nextHop != "" {
		node.SendMessage(msg, nextHop)
	}
}

func (p *SynapseProtocol) handleFound(msg *Message, node *Node) {
	// Update routing information
	node.Routing.UpdateRoute(msg.Key, msg.Source)

	// Handle based on original operation
	if msg.Type == "GET" {
		// Return value to requester
		fmt.Printf("Value found for key %s: %v\n", msg.Key, msg.Value)
	} else if msg.Type == "PUT" && msg.MRR > 0 {
		// Replicate to other networks
		msg.MRR--
		p.HandleOPE(msg, node)
	}
}

func (p *SynapseProtocol) handleJoin(msg *Message, node *Node) {
	// Handle join request
	fmt.Printf("Node %s joining network\n", msg.Source)
}
