package polyphonic

import (
	"encoding/json"
	"maps"
	"os"
	"sync"

	"github.com/yukumo-group/Chinese2KanaConverter/internal/cpyconverter"
)

// Manager manages polyphonic phrases
type Manager struct {
	sync.RWMutex
	heteronym      map[string]string
	targetFilePath string
}

// NewManager creates new manager
func NewManager() *Manager {
	return &Manager{
		heteronym:      make(map[string]string),
		targetFilePath: "polyphonic.json",
	}
}

// NewManagerFromFile creates new manager through reading .json file
func NewManagerFromFile(
	targetFilePath string,
) (*Manager, error) {
	newManager := Manager{}
	data, err := os.ReadFile(targetFilePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &newManager)
	return &newManager, err
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
