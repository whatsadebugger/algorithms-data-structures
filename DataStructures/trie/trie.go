package trie

// AlphabetSize number of letters
const AsciiValues = 128

// Node our only kind of node in the trie
type Node struct {
	children  []*Node
	endOfWord bool
}

// NewNode creates a new node initialized with a slice of pointers to child nodes
func NewNode() *Node {
	return &Node{
		children: make([]*Node, AsciiValues),
	}
}

// Search returns true if a key exists
func (root *Node) Search(key string) bool {
	current := root
	for _, c := range key {
		if current.children[c] == nil {
			return false
		}
		current = current.children[c]
	}
	return current.endOfWord
}

// Insert allows you to insert any valid lowercase english word
func (root *Node) Insert(key string) {
	current := root
	for _, c := range key {
		if current.children[c] == nil {
			current.children[c] = NewNode()
		}
		current = current.children[c]
	}
	current.endOfWord = true
}
