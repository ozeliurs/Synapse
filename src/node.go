package synapse

import (
	"fmt"
	"net"
	"sync"
)

type Node struct {
	IP         string
	Port       int
	Networks   []*Network
	TagManager *TagManager
	Protocol   *SynapseProtocol
	Routing    *RoutingTable
	listener   net.Listener
	mu         sync.Mutex
}

func NewNode(ip string, port int) *Node {
	return &Node{
		IP:         ip,
		Port:       port,
		Networks:   make([]*Network, 0),
		TagManager: NewTagManager(),
		Protocol:   NewSynapseProtocol(),
		Routing:    NewRoutingTable(),
	}
}

func (n *Node) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", n.IP, n.Port))
	if err != nil {
		return err
	}
	n.listener = listener

	go n.listen()
	return nil
}

func (n *Node) listen() {
	for {
		conn, err := n.listener.Accept()
		if err != nil {
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	message := &Message{}
	if err := message.Decode(conn); err != nil {
		return
	}

	n.Protocol.HandleMessage(message, n)
}

func (n *Node) OPE(code string, key string, value interface{}) {
	tag := n.TagManager.NewTag(n.IP)
	msg := NewMessage(code, key, value, n.IP, "")
	msg.Tag = tag

	n.Protocol.HandleOPE(msg, n)
}

func (n *Node) SendMessage(msg *Message, targetAddr string) error {
	conn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	return msg.Encode(conn)
}

func (n *Node) JoinNetwork(network *Network) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Networks = append(n.Networks, network)
	network.AddNode(n)
}

func (n *Node) GetAddress() string {
	return fmt.Sprintf("%s:%d", n.IP, n.Port)
}
