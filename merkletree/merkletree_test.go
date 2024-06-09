package merkletree

import (
	"bytes"
	"testing"

	"github.com/safwentrabelsi/merkletree-implementation/utils"
)

func TestMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := NewMerkleTree(data, utils.SHA256Hash)
	rootHash := tree.Root.Hash

	proof, err := tree.GenerateMerkleProof([]byte("data1"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid := VerifyMerkleProof(rootHash, []byte("data1"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for data1")
	}

	proof2, err := tree.GenerateMerkleProof([]byte("data2"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid2 := VerifyMerkleProof(rootHash, []byte("data2"), proof2, utils.SHA256Hash)
	if !isValid2 {
		t.Fatalf("Proof verification failed for data2")
	}

	// Test insert leaf
	tree.InsertLeaf([]byte("data5"))
	newRootHash := tree.Root.Hash
	if bytes.Equal(rootHash, newRootHash) {
		t.Fatalf("Root hash should change after inserting a new leaf")
	}

	proof3, err := tree.GenerateMerkleProof([]byte("data5"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid3 := VerifyMerkleProof(newRootHash, []byte("data5"), proof3, utils.SHA256Hash)
	if !isValid3 {
		t.Fatalf("Proof verification failed for data5")
	}

	// Test insert leaf
	tree.InsertLeaf([]byte("data6"))
	newRootHash2 := tree.Root.Hash
	if bytes.Equal(newRootHash, newRootHash2) {
		t.Fatalf("Root hash should change after inserting a new leaf")
	}

	proof4, err := tree.GenerateMerkleProof([]byte("data6"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid4 := VerifyMerkleProof(newRootHash2, []byte("data6"), proof4, utils.SHA256Hash)
	if !isValid4 {
		t.Fatalf("Proof verification failed for data6")
	}

	// Test update leaf
	err = tree.UpdateLeaf([]byte("data5"), []byte("newData5"))
	if err != nil {
		t.Fatalf("Failed to update leaf: %v", err)
	}
	updatedRootHash := tree.Root.Hash
	if bytes.Equal(newRootHash, updatedRootHash) {
		t.Fatalf("Root hash should change after updating a leaf")
	}
}
