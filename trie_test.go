package trie

import (
	"testing"
)

func TestAddSearch(t *testing.T) {
	type node struct {
		key   string
		value interface{}
	}
	type result struct {
		value  interface{}
		exists bool
	}
	tests := []struct {
		data     []node
		target   string
		expected result
	}{
		{[]node{}, "", result{nil, false}},
		{[]node{{"", nil}}, "", result{nil, false}},
		{[]node{}, "bike", result{nil, false}},
		{[]node{{"car", "pajero"}, {"card", "spade"}, {"care", "verb"}}, "car", result{"pajero", true}},
		{[]node{{"car", "pajero"}, {"card", "spade"}, {"care", "verb"}}, "card", result{"spade", true}},
		{[]node{{"car", "pajero"}, {"card", "spade"}, {"care", "verb"}}, "care", result{"verb", true}},
		{[]node{{"car", "pajero"}, {"card", "spade"}, {"care", "verb"}}, "vehicle", result{"pajero", false}},
	}
	for i, test := range tests {
		trie := NewTrie()
		for _, v := range test.data {
			trie.Add([]rune(v.key), v.value)
		}
		actual, exists := trie.Search([]rune(test.target))
		if exists != test.expected.exists {
			t.Errorf("%d, Error expected %t, got %t when searching for %s", i, test.expected.exists, exists, test.target)
		}
		if exists && actual != test.expected.value {
			t.Errorf("%d, Error expected %s, got %s", i, test.expected.value, actual)
		}
	}
}
