package trie_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/whatsadebugger/algorithms-data-structures/datastructures/trie"
)

func TestTrie(t *testing.T) {
	tree := trie.NewNode()

	tree.Insert("wow")
	if !tree.Search("wow") {
		t.Fatal("Could not find wow")
	}
	tree.Insert("you'd")
	if !tree.Search("you'd") {
		t.Fatal("Could not find you'd")
	}
}

func BenchmarkTrieInsert(b *testing.B) {
	b.StopTimer()

	words := make([]string, 0)
	tree := trie.NewNode()
	f, err := os.Open("wordsEn.txt")
	no(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Insert(word)
		}
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	b.StopTimer()

	words := make([]string, 0)
	tree := trie.NewNode()
	f, err := os.Open("../wordsEn.txt")
	no(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, word := range words {
		tree.Insert(word)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Search(word)
		}
	}
}

func no(err error) {
	if err != nil {
		panic(err)
	}
}
