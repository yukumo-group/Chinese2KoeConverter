package cpyconverter

import (
	"sync"

	"github.com/go-ego/gpy/phrase"
)

// loadOnce prevents the problem of concurrent map setting
var loadOnce sync.Once

// DumpHeteronymMap dumps map of heternym to the converter
func DumpHeteronymMap(
	heteronymMap map[string]string,
) {
	dumpFunc := func() {
		for chineseText, pinyin := range heteronymMap {
			phrase.DictAdd[chineseText] = pinyin
		}
	}
	loadOnce.Do(
		dumpFunc,
	)
}
