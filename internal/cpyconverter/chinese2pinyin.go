package cpyconverter

import (
	"strings"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
)

// ToPinyin converts chinese to pinyin.
// Since tone is not important it will remove tones
func ToPinyin(
	chineseText string,
	useHeteronym bool,
) []string {
	arg := gpy.Args{
		Style: gpy.Normal,
	}
	resultWithTone := phrase.Pinyin(
		chineseText,
	)
	resultNoTone := []string{}
	for _, pyWithTone := range resultWithTone {
		pyWithNoTone := strings.TrimSpace(
			gpy.ToFixed(
				pyWithTone,
				arg,
			),
		)
		resultNoTone = append(resultNoTone, pyWithNoTone)
	}
	return resultNoTone
}
