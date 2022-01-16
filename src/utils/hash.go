package utils

import "hash/fnv"

func Hash(str string) int {
	// Create a new hash.Hash32 computing the FNV-1 hash.
	h := fnv.New32a()
	// Write the bytes of the string to the hash.
	h.Write([]byte(str))
	// Sum the bytes and return the result.
	return int(h.Sum32())
}