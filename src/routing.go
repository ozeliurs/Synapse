package synapse

import "sync"

type RoutingTable struct {
	routes map[string]string // key -> nodeAddress
	mu     sync.RWMutex
}

func NewRoutingTable() *RoutingTable {
	return &RoutingTable{
		routes: make(map[string]string),
	}
}

func (rt *RoutingTable) UpdateRoute(key string, nodeAddr string) {
	rt.mu.Lock()
	defer rt.mu.Unlock()

	rt.routes[key] = nodeAddr
}

func (rt *RoutingTable) GetNextHop(key string) string {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	return rt.routes[key]
}

func (rt *RoutingTable) IsResponsible(key string) bool {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	_, exists := rt.routes[key]
	return exists
}
