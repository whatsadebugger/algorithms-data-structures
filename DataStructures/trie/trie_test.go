package trie_test

import (
	"github.com/whatsadebugger/algorithms-data-structures/DataStructures/trie"
	"testing"
)

func TestInsert(t *testing.T) {

}

func TestSearch(t *testing.T) {

}

func TestTree(t *testing.T) {
	keys := []string{
		"the", "a", "there",
		"answer", "any", "by",
		"bye", "their",
	}

	root := trie.Node{}

	// Construct trie
	for i := 0; i < len(keys); i++ {
		root.Insert(keys[i])
	}

	// Search for different keys
	if !root.Search("the") {
		t.fai
	}

	root.Search("these")
}
