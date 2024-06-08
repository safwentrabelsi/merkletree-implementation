package merkletree

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Hash   []byte
	Data   []byte
	Index  int
}

type MerkleTree struct {
	Root   *Node
	Leaves []*Node
}

type ProofNode struct {
	Hash []byte
	Left bool
}

func NewMerkleTree(data [][]byte) *MerkleTree {

	leaves := make([]*Node, len(data))
	for i, d := range data {
		hash := hash256(d)
		leaves[i] = &Node{Hash: hash, Data: d, Index: i}
	}

	tree := &MerkleTree{
		Leaves: leaves,
	}
	tree.Root = buildTree(leaves)
	return tree

}

func buildTree(nodes []*Node) *Node {

	if len(nodes) == 1 {
		return nodes[0]
	}

	var ParentNodes []*Node

	for i := 0; i < len(nodes); i += 2 {
		left := nodes[i]
		right := left

		if i+1 < len(nodes) {
			right = nodes[i+1]
		}

		parentHash := hash256(append(left.Hash, right.Hash...))
		parentNode := &Node{Hash: parentHash, Left: left, Right: right}
		left.Parent = parentNode
		right.Parent = parentNode
		ParentNodes = append(ParentNodes, parentNode)
	}

	return buildTree(ParentNodes)

}

func (m *MerkleTree) GenerateMerkleProof(data []byte) ([]ProofNode, error) {
	var leaf *Node
	for _, l := range m.Leaves {
		if bytes.Equal(l.Data, data) {
			leaf = l
		}
	}

	if leaf == nil {
		return nil, fmt.Errorf("data not found in tree")
	}

	var proof []ProofNode
	current := leaf
	for current.Parent != nil {

		sibling, left := current.getSibling()

		if sibling != nil {
			proof = append(proof, ProofNode{Hash: sibling.Hash, Left: left})

		}
		current = current.Parent
	}

	return proof, nil
}

func (n *Node) getSibling() (*Node, bool) {
	if n.Parent == nil {
		return nil, false
	}

	if n.Parent.Left == n {
		return n.Parent.Right, false
	}
	return n.Parent.Left, true
}

func VerifyMerkleProof(rootHash []byte, data []byte, proof []ProofNode) bool {

	currentHash := hash256(data)

	for _, p := range proof {
		if p.Left {
			currentHash = hash256(append(p.Hash, currentHash...))

		} else {
			currentHash = hash256(append(currentHash, p.Hash...))

		}
	}
	return bytes.Equal(currentHash, rootHash)
}

func hash256(d []byte) []byte {
	hash := sha256.Sum256(d)
	return hash[:]
}
