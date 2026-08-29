package cpyconverter

import (
	"github.com/go-ego/gpy/phrase"
)

// DumHeteronymMap dumps map of heternym to the converter
func DumpHeteronymMap(
	heteronymMap map[string]string,
) {
	for chineseText, pinyin := range heteronymMap {
		phrase.DictAdd[chineseText] = pinyin
	}
}
