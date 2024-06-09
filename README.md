# Merkle Tree Implementation

This project provides an implementation of a Merkle Tree in Go. It supports creating a Merkle Tree, generating Merkle proofs, verifying proofs, and inserting and updating leaves.

## Key Features

- **Incremental Updates**: The Merkle tree is updated incrementally when an existing leaf is updated. This ensures efficient updates without needing to rebuild the entire tree.
- **Handling Odd Number of Leaves**: When the number of leaves is odd, the implementation duplicates the last leaf to maintain a balanced tree structure. This choice simplifies the tree operations and ensures consistent performance.
- **No Sorting of Leaves**: The leaves are not sorted before building the tree, preserving the order of insertion.

## Prerequisites

Ensure you have Go 1.21.3 installed on your system.

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/safwentrabelsi/merkletree-implementation.git
cd merkletree-implementation
```

## Building the Application

You can build the application using the provided Makefile. This will compile the Go code and produce a binary in the `build` folder.

```bash
make build
```

## Running Tests

To run the tests, use the following command. This will execute all tests and provide a coverage report.

```bash
make test
```

## Running the Application

After building the application, you can run it using the following command:

```bash
make run
```

This will build the application (if not already built) and then run it.

## Cleaning Up

To clean up the build artifacts and cached test results, use the following command:

```bash
make clean
```

## Usage Example

Here is an example of how to use the Merkle Tree implementation in your Go code:

```go
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

    // Generate proof for a leaf
    proof, err := tree.GenerateMerkleProof([]byte("data1"))
    if err != nil {
        fmt.Println("Error generating proof:", err)
        return
    }
    fmt.Println("Proof for data1 generated successfully")

    // Verify the proof
    isValid := merkletree.VerifyMerkleProof(tree.Root.Hash, []byte("data1"), proof, utils.SHA256Hash)
    fmt.Printf("Proof valid: %v\n", isValid)

    // Insert new leaf
    tree.InsertLeaf([]byte("data5"))
    fmt.Printf("New root hash after inserting data5: %x\n", tree.Root.Hash)
}
```
