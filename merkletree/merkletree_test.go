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

	// Test proof generation and verification for data1
	proof, err := tree.GenerateMerkleProof([]byte("data1"))
	if err != nil {
		t.Fatalf("Failed to generate proof for data1: %v", err)
	}
	isValid := VerifyMerkleProof(rootHash, []byte("data1"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for data1")
	}

	// Test proof generation and verification for data2
	proof, err = tree.GenerateMerkleProof([]byte("data2"))
	if err != nil {
		t.Fatalf("Failed to generate proof for data2: %v", err)
	}
	isValid = VerifyMerkleProof(rootHash, []byte("data2"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for data2")
	}

	// Test proof generation for non-existent data
	proof, err = tree.GenerateMerkleProof([]byte("nonExistentData"))
	if err == nil {
		t.Fatalf("Expected error for non-existent data proof generation")
	}

	// Test insert leaf
	tree.InsertLeaf([]byte("data5"))
	newRootHash := tree.Root.Hash
	if bytes.Equal(rootHash, newRootHash) {
		t.Fatalf("Root hash should change after inserting a new leaf")
	}

	proof, err = tree.GenerateMerkleProof([]byte("data5"))
	if err != nil {
		t.Fatalf("Failed to generate proof for data5: %v", err)
	}
	isValid = VerifyMerkleProof(newRootHash, []byte("data5"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for data5")
	}

	// Test insert another leaf
	tree.InsertLeaf([]byte("data6"))
	newRootHash2 := tree.Root.Hash
	if bytes.Equal(newRootHash, newRootHash2) {
		t.Fatalf("Root hash should change after inserting a new leaf")
	}

	proof, err = tree.GenerateMerkleProof([]byte("data6"))
	if err != nil {
		t.Fatalf("Failed to generate proof for data6: %v", err)
	}
	isValid = VerifyMerkleProof(newRootHash2, []byte("data6"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for data6")
	}

	// Test update leaf
	err = tree.UpdateLeaf([]byte("data5"), []byte("newData5"))
	if err != nil {
		t.Fatalf("Failed to update leaf: %v", err)
	}
	updatedRootHash := tree.Root.Hash
	if bytes.Equal(newRootHash2, updatedRootHash) {
		t.Fatalf("Root hash should change after updating a leaf")
	}

	// Verify data5 doesn't is replaced
	proof, err = tree.GenerateMerkleProof([]byte("data5"))
	if err == nil || err.Error() != "data not found in tree" {
		t.Fatal("Generate proof should return error")
	}

	proof, err = tree.GenerateMerkleProof([]byte("newData5"))
	if err != nil {
		t.Fatalf("Failed to generate proof for newData5: %v", err)
	}

	isValid = VerifyMerkleProof(updatedRootHash, []byte("newData5"), proof, utils.SHA256Hash)
	if !isValid {
		t.Fatalf("Proof verification failed for newData5")
	}
}
