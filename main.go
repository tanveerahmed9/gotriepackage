package mytrie

import (
	"fmt"

	"unicode"
)

const AlphabetSize = 26

// make structure for node and trie

// initiate function which add a root structure

// insert method

// search method

// Node struct which has children Node pointer and the isEndofWord boolean which represent whether the character is end of any word

type Node struct {
	children [AlphabetSize]*Node

	isEndofWord bool
}

// trie struct represents a trie that is a pointer to the root node

type trie struct {
	root *Node
}

// insert will take a word and add it to the trie

func (t *trie) insert(w string) {

	currentNode := t.root

	for i := 0; i < len(w); i++ {

		charIndex := unicode.ToLower(rune(w[i])) - 'a' // convert any upper case entry to avoid out of bound and index mismatch

		if currentNode.children[charIndex] == nil {

			currentNode.children[charIndex] = &Node{}

		}

		currentNode = currentNode.children[charIndex]

	}

	currentNode.isEndofWord = true // after the word insertion ends

}

// Search method returns true if the word (w) in present in the trie

func (t *trie) Searchtrie(w string) bool {

	currentNode := t.root

	for i := 0; i < len(w); i++ {

		charIndex := unicode.ToLower(rune(w[i])) - 'a' // convert any upper case entry to avoid out of bound and index mismatch

		if currentNode.children[charIndex] == nil {

			return false

		}

		currentNode = currentNode.children[charIndex]

	}

	return currentNode.isEndofWord

}

func inittrie() *trie {

	result := &trie{root: &Node{}}

	return result

}

func main() {

	mytrie := inittrie()

	toAdd := []string{

		"Tbahmed",

		"tanveerahmed",

		"tbahmed9",

		"mytbahmed",

		"tbalternate",
	}

	for _, v := range toAdd {

		fmt.Println("Adding", v)

		mytrie.insert(v)

	}

}
