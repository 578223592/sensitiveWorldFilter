package wordsFilter

type Node struct {
	isEndNode bool
	character rune
	children  map[rune]*Node
}

type TreeRoot Node

// NewNode New creates a node.
func NewNode(isEndNode bool, character rune) *Node {
	return &Node{
		isEndNode: isEndNode,
		character: character,
		children:  make(map[rune]*Node),
	}
}

// Add sensitive words to specified sensitive words Map.
// make sure that only tree root call this function
func (root *TreeRoot) add(text string) {
	if text == "" {
		return
	}
	curNode := (*Node)(root)
	textRunes := []rune(text)
	for index := 0; index < len(textRunes); index++ {
		curRune := textRunes[index]
		if nextNode, ok := curNode.children[curRune]; ok {
			curNode = nextNode
		} else {
			newNode := NewNode(false, curRune)
			curNode.children[curRune] = newNode
			curNode = newNode
		}
		if index == len(textRunes)-1 {
			curNode.isEndNode = true
		}
	}
}

// Remove specified sensitive words from sensitive word map.
func (root *TreeRoot) remove(text string) {
	textRunes := []rune(text)
	var curNode = (*Node)(root)
	for index := 0; index < len(textRunes); index++ {
		if nextNode, ok := curNode.children[textRunes[index]]; ok {
			curNode = nextNode
		} else {
			return //不包含这个单词
		}
		if index == len(textRunes)-1 {
			curNode.isEndNode = false //soft delete
		}
	}
}

// Whether the string contains sensitive words.
// 找的思路就是往下沿着text遍历整个tree，如果发现了走到了结尾，就可以标记了，
// 如果走的过程中发现了树中没有对应的节点，那么就重新开始匹配
func (root *TreeRoot) contains(text string) (bool, string) {
	textRune := []rune(text)
	curParent := (*Node)(root)
	var sensitiveWordStartIndex = 0
	for index := 0; index < len(textRune); index++ {
		var curRune = textRune[index]
		curNode, ok := curParent.children[curRune]
		if !ok || (!curNode.isEndNode && index == len(textRune)-1) {
			// 重新开始匹配
			curParent = (*Node)(root)
			index = sensitiveWordStartIndex // index在本次循环结束会自动++，因此这里不用++
			sensitiveWordStartIndex++
			continue
		}

		if curNode.isEndNode {
			return true, string(textRune[sensitiveWordStartIndex : index+1])
		}

		curParent = curNode
	}
	return false, ""
}

func (root *TreeRoot) filterAll(text string) string {
	var resultRunes = make([]rune, 0)

	textRune := []rune(text)
	curParent := (*Node)(root)
	var sensitiveWordStartIndex = 0
	for index := 0; index < len(textRune); index++ {
		var curRune = textRune[index]
		curNode, found := curParent.children[curRune]
		if !found || (!curNode.isEndNode && index == len(textRune)-1) {
			// 重新开始匹配
			resultRunes = append(resultRunes, curRune)
			curParent = (*Node)(root)
			index = sensitiveWordStartIndex // index在本次循环结束会自动++，因此这里不用++
			sensitiveWordStartIndex++
			continue
		}

		if curNode.isEndNode {
			//匹配上了，这中间的都不要了
			sensitiveWordStartIndex = index + 1
			curParent = (*Node)(root)
		} else { //往下遍历
			curParent = curNode
		}

	}
	return string(resultRunes)
}
