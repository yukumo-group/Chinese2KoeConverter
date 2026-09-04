package polyphonic

import (
	"maps"
	"sync"

	"github.com/yukumo-group/Chinese2KanaConverter/internal/cpyconverter"
)

var securedDumpLoad sync.RWMutex

// LoadPolyphonics directly loads the polyphonics
func LoadPolyphonics(
	polyphonics map[string]string,
) {
	cpyconverter.DumpHeteronymMap(
		maps.Clone(polyphonics),
	)
}

// SafeLoadPolyphonics loads the the polyphonics safely.
func SafeLoadPolyphonics(
	polyphonics map[string]string,
) {
	securedDumpLoad.Lock()
	defer securedDumpLoad.Unlock()
	cpyconverter.DumpHeteronymMap(
		maps.Clone(polyphonics),
	)
}
