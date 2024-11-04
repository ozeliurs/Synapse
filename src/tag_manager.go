package synapse

import (
	"fmt"
	"sync"
	"time"
)

type TagManager struct {
	tags map[string]bool
	mu   sync.RWMutex
}

func NewTagManager() *TagManager {
	return &TagManager{
		tags: make(map[string]bool),
	}
}

func (tm *TagManager) NewTag(source string) string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tag := fmt.Sprintf("%s-%d", source, time.Now().UnixNano())
	tm.tags[tag] = true
	return tag
}

func (tm *TagManager) IsProcessed(tag string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	return tm.tags[tag]
}

func (tm *TagManager) PushTag(tag string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tags[tag] = true
}
