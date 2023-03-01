package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"crypto/sha256"
)

func main() {
	// Open the input file
	file, err := os.Open("transactions.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the transaction hashes
	var hashes [][]byte

	// Read the transaction hashes from the file
	for scanner.Scan() {
		// Decode the hex-encoded hash into a byte slice
		hash, err := hex.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}

		// Add the hash to the slice
		hashes = append(hashes, hash)
	}

	// Compute the Merkle root
	root := computeMerkleRoot(hashes)

	// Print the root
	fmt.Printf("Merkle root: %x\n", root)
}

func computeMerkleRoot(hashes [][]byte) []byte {
	// Base case: if there's only one hash, return it
	if len(hashes) == 1 {
		return hashes[0]
	}

	// If the number of hashes is odd, duplicate the last hash
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// Create a new slice to hold the next level of hashes
	nextLevel := make([][]byte, len(hashes)/2)

	// Hash each pair of hashes and add the result to the next level slice
	for i := 0; i < len(hashes); i += 2 {
		hash := sha256.Sum256(append(hashes[i], hashes[i+1]...))
		nextLevel[i/2] = hash[:]
	}

	// Recursively compute the Merkle root of the next level
	return computeMerkleRoot(nextLevel)
}
