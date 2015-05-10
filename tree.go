package merkletree

import (
	"fmt"
	"strings"
	"crypto/sha256"
	"encoding/hex"
)

var (
	SHA256 = func(s string) string {
		hasher := sha256.New()
		hasher.Write([]byte(s))
		return hex.EncodeToString(hasher.Sum(nil))
	}

	DOUBLE_SHA256 = func(s string) string {
		return SHA256(SHA256(s))
	}
)

type HashFn func(s string) string

type Tree struct {
	Left  *Tree
	Right *Tree
	Hash  string
}

func (t *Tree) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

func (t * Tree) PrintTree(level int) {
	// Indentations and Prefix
	prefix := ""
	indentation := ""
	if level > 0 {
		prefix = "|--"
		indentation = strings.Repeat("\t", level)
	}
	fmt.Println(indentation, prefix, t.Hash)

	// Print children trees if there is any
	if t.IsLeaf() == false {
		t.Left.PrintTree(level + 1)
		t.Right.PrintTree(level + 1)
	}
}

func Hash(hashes []string, hasher HashFn) string {
	tree := BuildTree(hashes, hasher)
	return tree.Hash
}

func BuildTree(hashes []string, hasher HashFn) Tree {

	// Build first level of trees
	trees := []Tree{}
	for _, hash := range hashes {
		t := Tree{
			Hash: hash,
		}
		trees = append(trees, t)
	}

	parentTrees := hashTrees(trees, hasher)
	for {
		if len(parentTrees) == 1 {
			break
		}
		parentTrees = hashTrees(parentTrees, hasher)
	}

	rootTree := parentTrees[0]

	return rootTree
}

func hashTrees(trees []Tree, hasher HashFn) []Tree {
	if len(trees) %2 != 0 {
		trees = append(trees, trees[len(trees) - 1])
	}

	parents := []Tree{}
	for i := 0; i < len(trees); i+=2 {
		t := Tree {
			Left: &trees[i],
			Right: &trees[i+1],
			Hash: concatAndHash(trees[i].Hash, trees[i+1].Hash, hasher),
		}
		parents = append(parents, t)
	}

	return parents
}

func concatAndHash(a string, b string, hasher HashFn) string {
	concat := a + b
	// Perform hash
	hash := hasher(concat)
	return hash
}