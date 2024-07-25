package wordsFilter

import (
	"bytes"

	"strings"
	"sync"
)

type WordsFilter struct {
	treeRoot *TreeRoot
	mutex    sync.RWMutex
}

// NewWordsFilter
//
//	@Description: New creates a words filter.
//	@param ignoreSpace
//	@return *WordsFilter
func NewWordsFilter() *WordsFilter {
	return &WordsFilter{
		treeRoot: (*TreeRoot)(NewNode(false, 1)),
	}
}

// Add  sensitive text lists into sensitive word tree nodes
func (wf *WordsFilter) Add(texts ...string) {
	wf.mutex.Lock()
	defer wf.mutex.Unlock()
	for _, text := range texts {
		wf.treeRoot.add(text)
	}

}

func (wf *WordsFilter) IsContainsSensitiveWord(text string) (bool, string) {
	if len(text) == 0 {
		return false, ""
	}
	wf.mutex.RLock()
	defer wf.mutex.RUnlock()
	return wf.treeRoot.contains(text)
}

// RemoveSensitiveWords
//
//	@Description: 移除敏感词
//	@receiver wf
//	@param texts
func (wf *WordsFilter) RemoveSensitiveWords(texts ...string) {
	for _, text := range texts {
		wf.mutex.Lock()
		wf.treeRoot.remove(text)
		wf.mutex.Unlock()
	}
}

// Strip space
func stripSpace(str string) string {
	fields := strings.Fields(str)
	var bf bytes.Buffer
	for _, field := range fields {
		bf.WriteString(field)
	}
	return bf.String()
}

func (wf *WordsFilter) FilterAll(text string) string {
	newText := wf.treeRoot.filterAll(text)
	return newText
}
