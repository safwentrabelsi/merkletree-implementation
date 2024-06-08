package merkletree

import (
	"testing"
)

func TestMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := NewMerkleTree(data)
	rootHash := tree.Root.Hash

	proof, err := tree.GenerateMerkleProof([]byte("data1"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid := VerifyMerkleProof(rootHash, []byte("data1"), proof)
	if !isValid {
		t.Fatalf("Proof verification failed for data1")
	}

	proof2, err := tree.GenerateMerkleProof([]byte("data2"))
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}

	isValid2 := VerifyMerkleProof(rootHash, []byte("data2"), proof2)
	if !isValid2 {
		t.Fatalf("Proof verification failed for data2")
	}
}
