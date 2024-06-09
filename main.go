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
		fmt.Println("List of proofs for data1:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}

	isValid := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data1"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid: %v\n", isValid)

	// right example
	proof, err = tree.GenerateMerkleProof([]byte("data2"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for data2:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}

	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data2"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for data2: %v\n", isValid)

	// Insert new leaves
	tree.InsertLeaf([]byte("data5"))
	tree.InsertLeaf([]byte("data6"))

	proof, err = tree.GenerateMerkleProof([]byte("data6"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for data6:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data6"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for data6: %v\n", isValid)

	tree.InsertLeaf([]byte("data7"))

	proof, err = tree.GenerateMerkleProof([]byte("data7"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for data7:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data7"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for data7: %v\n", isValid)

	tree.InsertLeaf([]byte("data8"))

	proof, err = tree.GenerateMerkleProof([]byte("data8"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for data8:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data8"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for data8: %v\n", isValid)

	tree.InsertLeaf([]byte("data9"))

	proof, err = tree.GenerateMerkleProof([]byte("data9"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for data9:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data9"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for data9: %v\n", isValid)

	// Test update leaf
	err = tree.UpdateLeaf([]byte("data5"), []byte("newData5"))
	if err != nil {
		fmt.Println(err)
	}
	proof, err = tree.GenerateMerkleProof([]byte("newData5"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("List of proofs for newData5:")
		for _, node := range proof {
			fmt.Printf("Hash: %x, Left: %v\n", node.Hash, node.Left)
		}
	}
	fmt.Printf("Root: %x\n", tree.Root.Hash)
	isValid = merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("newData5"), proof, utils.SHA256Hash)
	fmt.Printf("Proof valid for newData5: %v\n", isValid)
}
