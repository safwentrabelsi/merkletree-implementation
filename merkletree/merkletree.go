package merkletree

import (
	"bytes"
	"fmt"
	"math"
)

// HashFunction defines a function type for hashing data
type HashFunction func(data []byte) []byte

// Node represents a node in the Merkle tree
type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Hash   []byte
	Data   []byte
	Index  int
}

// MerkleTree represents the Merkle tree structure
type MerkleTree struct {
	Root     *Node
	Leaves   []*Node
	hashFunc HashFunction
}

// ProofNode represents a node in the Merkle proof
type ProofNode struct {
	Hash []byte
	Left bool
}

// NewMerkleTree creates a new Merkle tree from the provided data and hash function
func NewMerkleTree(data [][]byte, hashFunc HashFunction) *MerkleTree {
	leaves := make([]*Node, len(data))
	for i, d := range data {
		hash := hashFunc(d)
		leaves[i] = &Node{Hash: hash, Data: d, Index: i}
	}

	tree := &MerkleTree{
		Leaves:   leaves,
		hashFunc: hashFunc,
	}
	tree.Root = buildTree(leaves, hashFunc)
	return tree
}

// GenerateMerkleProof generates a Merkle proof for the given data
func (m *MerkleTree) GenerateMerkleProof(data []byte) ([]ProofNode, error) {
	leaf, err := m.findLeaf(data)

	if err != nil {
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

// VerifyMerkleProof verifies the given Merkle proof against the root hash and data
func VerifyMerkleProof(rootHash []byte, data []byte, proof []ProofNode, hashFunc HashFunction) bool {
	currentHash := hashFunc(data)

	for _, p := range proof {
		if p.Left {
			currentHash = hashFunc(append(p.Hash, currentHash...))
		} else {
			currentHash = hashFunc(append(currentHash, p.Hash...))
		}
	}
	return bytes.Equal(currentHash, rootHash)
}

// InsertLeaf inserts a new leaf into the Merkle tree
func (m *MerkleTree) InsertLeaf(data []byte) {
	newLeaf := &Node{Data: data, Hash: m.hashFunc(data), Index: len(m.Leaves)}
	m.Leaves = append(m.Leaves, newLeaf)

	if len(m.Leaves)%2 == 1 {
		m.Root = buildTree(m.Leaves, m.hashFunc)
	} else {
		parent := m.Leaves[len(m.Leaves)-2].Parent
		parent.Right = newLeaf
		oldHash := parent.Hash
		parent.Hash = m.hashFunc(append(parent.Left.Hash, parent.Right.Hash...))
		newLeaf.Parent = parent
		m.incrementalUpdate(parent, oldHash)
	}
}

// UpdateLeaf updates the data of an existing leaf in the Merkle tree
func (m *MerkleTree) UpdateLeaf(oldData, newData []byte) error {
	leaf, err := m.findLeaf(oldData)

	if err != nil {
		return err
	}

	leaf.Data = newData
	oldHash := leaf.Hash
	leaf.Hash = m.hashFunc(newData)
	m.incrementalUpdate(leaf, oldHash)
	return nil
}

// buildTree recursively builds the Merkle tree from the leaves
func buildTree(nodes []*Node, hashFunc HashFunction) *Node {
	if len(nodes) == 1 {
		return nodes[0]
	}

	var parentNodes []*Node

	for i := 0; i < len(nodes); i += 2 {
		left := nodes[i]
		right := left

		if i+1 < len(nodes) {
			right = nodes[i+1]
		}

		parentHash := hashFunc(append(left.Hash, right.Hash...))
		parentNode := &Node{Hash: parentHash, Left: left, Right: right}
		left.Parent = parentNode
		right.Parent = parentNode
		parentNodes = append(parentNodes, parentNode)
	}

	return buildTree(parentNodes, hashFunc)
}

// incrementalUpdate updates the tree from the given node up to the root
func (m *MerkleTree) incrementalUpdate(node *Node, oldHash []byte) {
	current := node
	currentOldHash := oldHash
	for current.Parent != nil {
		sibling, left := current.getSibling()
		if bytes.Equal(sibling.Hash, currentOldHash) {
			currentOldHash = current.Parent.Hash
			current.Parent.Hash = m.hashFunc(append(current.Hash, current.Hash...))
		} else {
			if !left {
				current.Parent.Hash = m.hashFunc(append(current.Hash, sibling.Hash...))
			} else {
				current.Parent.Hash = m.hashFunc(append(sibling.Hash, current.Hash...))
			}
		}
		current = current.Parent
	}
	m.Root = current
}

// findLeaf searches for leaf of the provided data
func (m *MerkleTree) findLeaf(d []byte) (*Node, error) {
	var leaf *Node

	for _, l := range m.Leaves {
		if bytes.Equal(l.Data, d) {
			leaf = l
			break
		}
	}

	if leaf == nil {
		return nil, fmt.Errorf("data not found in tree")
	}
	return leaf, nil
}

// calculateDepth calculates the depth of the tree based on the number of leaves
func calculateDepth(numLeaves int) int {
	return int(math.Ceil(math.Log2(float64(numLeaves))))
}
