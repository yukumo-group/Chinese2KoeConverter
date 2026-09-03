package language

// Chunk defines the chunked text
type Chunk struct {
	Text      string
	IsChinese bool
}

// SeparateToChunks separate text into chunks
func SeparateToChunks(
	text string,
) []*Chunk {
	result := []*Chunk{}
	runedText := []rune(text)
	currentChunk := ""
	currentIsChineseState := true
	for i := 0; i < len(runedText); i++ {
		newIsChineseState := IsChinese(string(runedText[i]))
		if i == 0 {
			currentIsChineseState = newIsChineseState
			currentChunk = string(runedText[i])
		} else {
			if currentIsChineseState == newIsChineseState {
				currentChunk += string(runedText[i])
			} else {
				tmpResult := &Chunk{
					Text:      currentChunk,
					IsChinese: currentIsChineseState,
				}
				result = append(result, tmpResult)
				currentIsChineseState = newIsChineseState
				currentChunk = string(runedText[i])
			}
		}
	}
	if currentChunk != "" {
		tmpResult := &Chunk{
			Text:      currentChunk,
			IsChinese: currentIsChineseState,
		}
		result = append(result, tmpResult)
	}
	return result
}
