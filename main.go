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

	tree.InsertLeaf([]byte("data5"))
	tree.InsertLeaf([]byte("data6"))

	proof3, err := tree.GenerateMerkleProof([]byte("data6"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof3 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid3 := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data6"), proof3, utils.SHA256Hash)
	fmt.Printf("Proof valid for data6: %v\n", isValid3)

	tree.InsertLeaf([]byte("data7"))

	proof4, err := tree.GenerateMerkleProof([]byte("data7"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof4 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid4 := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data7"), proof4, utils.SHA256Hash)
	fmt.Printf("Proof valid for data7: %v\n", isValid4)

	tree.InsertLeaf([]byte("data8"))

	proof5, err := tree.GenerateMerkleProof([]byte("data8"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof5 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid5 := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data8"), proof5, utils.SHA256Hash)
	fmt.Printf("Proof valid for data8: %v\n", isValid5)

	tree.InsertLeaf([]byte("data9"))

	proof6, err := tree.GenerateMerkleProof([]byte("data9"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof6 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid6 := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data9"), proof6, utils.SHA256Hash)
	fmt.Printf("Proof valid for data9: %v\n", isValid6)

	// Test update leaf
	err = tree.UpdateLeaf([]byte("data5"), []byte("newData5"))
	if err != nil {
	}
	proof5, err = tree.GenerateMerkleProof([]byte("newData5"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs:")
		for _, node := range proof6 {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid5 = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data9"), proof5, utils.SHA256Hash)
	fmt.Printf("Proof valid for newdata5: %v\n", isValid5)

}
