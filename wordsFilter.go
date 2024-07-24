package wordsFilter

import (
	"bytes"

	"strings"
	"sync"
)

// defaultIgnoreRunes 对敏感词表操作对时候会忽略对字符
var DefaultIgnoreRunes = []rune{}

type WordsFilter struct {
	config   Config
	treeRoot *TreeRoot
	mutex    sync.RWMutex
}

type Config struct {
	StripSpace  bool
	IgnoreRunes []rune //todo 看下这个功能开启没哟
}

// New creates a words filter.
func NewWordsFilter(ignoreRunes []rune, ignoreSpace bool) *WordsFilter {
	return &WordsFilter{
		config: Config{
			StripSpace:  ignoreSpace,
			IgnoreRunes: ignoreRunes,
		},
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
	if wf.config.StripSpace {
		text = stripSpace(text)
	}
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
		if wf.config.StripSpace {
			text = stripSpace(text)
		}
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
