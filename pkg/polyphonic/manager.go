package polyphonic

import (
	"maps"
	"sync"

	"github.com/yukumo-group/Chinese2KanaConverter/internal/cpyconverter"
)

// Manager manages polyphonic phrases
type Manager struct {
	sync.RWMutex
	heteronym map[string]string
}

// NewManager creates new manager
func NewManager() *Manager {
	return &Manager{
		heteronym: map[string]string{},
	}
}

// load loads the heteronyms.
// **Unlocked!**
func (manager *Manager) load() {
	cpyconverter.DumpHeteronymMap(
		maps.Clone(manager.heteronym),
	)
}

// Initialize initializes the manager
func (manager *Manager) Initialize() {
	manager.RLock()
	defer manager.RUnlock()
	manager.load()
}
