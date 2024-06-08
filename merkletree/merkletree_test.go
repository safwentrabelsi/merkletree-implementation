package merkletree

import (
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
}
