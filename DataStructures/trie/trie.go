package trie

const alphabetSize = 26

// Node is an efficient structure for data retrieval.
// we will use it for searching for valid words.
type Node struct {
	children    *[alphabetSize]Node
	isEndOfWord bool
}

// Insert a node into the Trie
func (n *Node) Insert() {

}

// Search for a key in the Trie
func (n *Node) Search() {

}
