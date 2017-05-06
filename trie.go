package trie

// Trie is the node in an ordered tree data structure
type Trie struct {
	children map[rune]*Trie
	complete bool
	value    interface{}
}

// Add inserts a word and value into the trie tree
func (t *Trie) Add(word []rune, value interface{}) {
	if len(word) == 0 {
		return
	}
	c := word[0]
	if _, exists := t.children[c]; !exists {
		t.children[c] = NewTrie()
	}
	if len(word) == 1 {
		t.children[c].complete = true
		t.children[c].value = value
	} else {
		t.children[c].Add(word[1:], value)
	}
}

// Search finds and returns value, true if it exists or nil, false otherwise
func (t *Trie) Search(word []rune) (interface{}, bool) {
	if len(word) == 0 {
		return t.value, t.complete
	}
	c := word[0]
	if _, exists := t.children[c]; exists {
		return t.children[c].Search(word[1:])
	}
	return nil, false
}

// NewTrie creates a new trie node
func NewTrie() *Trie {
	return &Trie{children: make(map[rune]*Trie), complete: false}
}
