package main

import (
	"fmt"

	"github.com/safwentrabelsi/merkletree-implementation/merkletree"
	"github.com/safwentrabelsi/merkletree-implementation/utils"
)

func main() {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := merkletree.NewMerkleTree(data, utils.SHA256Hash)
	fmt.Printf("Root hash: %x\n", tree.Root.Hash)

	// left example
	proof, err := tree.GenerateMerkleProof([]byte("data1"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}

	isValid := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data1"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid: %v\n", isValid)

	// right example
	proof2, err := tree.GenerateMerkleProof([]byte("data2"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof2 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}

	isValid2 := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data2"), proof2, utils.SHA256Hash)
	fmt.Printf("Proof valid for data2: %v\n", isValid2)
}
