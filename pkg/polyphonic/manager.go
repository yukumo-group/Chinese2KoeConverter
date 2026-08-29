package polyphonic

import (
	"sync"
)

// Manager manages polyphonic phrases
type Manager struct {
	sync.RWMutex
}
