package language

import (
	"testing"

	"github.com/yukumo-group/Chinese2KanaConverter/internal/process"
)

// TestIsChinese tests if certain words are Chinese
func TestIsChinese(t *testing.T) {
	tests := []string{
		"aobcccc",
		"嘴巴",
		"嘴巴，最最最",
		"嘴巴ab",
		"〇",
	}
	expectedResult := []bool{
		false,
		true,
		false,
		false,
		true,
	}
	for i, val := range tests {
		result := IsChinese(val)
		if expectedResult[i] != result {
			t.Errorf(
				"Expected test result of %s to be %s, got %s",
				val,
				process.ConvertBoolToString(
					expectedResult[i],
				),
				process.ConvertBoolToString(
					result,
				),
			)
		}
	}
}
