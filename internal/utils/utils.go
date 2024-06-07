// Package utils provides utility functions for common operations.
//
// This package contains various utility functions that can be used
// across different parts of an application to perform common tasks
// such as removing duplicates from slices, string manipulations, etc.
package utils

// RemoveDuplicatedStringE removes duplicate strings from a slice.
//
// This function takes a slice of strings as input and returns a new slice
// containing only the unique strings from the original slice. The order
// of the strings in the resulting slice is not guaranteed to be the same
// as in the input slice.
//
// Example usage:
// input := []string{"apple", "banana", "apple", "orange", "banana"}
// result := RemoveDuplicatedStringE(input)
// fmt.Println(result) // Output might be: ["apple", "banana", "orange"]
func RemoveDuplicatedStringE(s []string) []string {
	var cleanedSlice []string
	var set map[string]int = make(map[string]int)

	// Save all items into a set to remove duplicates
	for idx, e := range s {
		set[e] = idx
	}

	// Save all unique items into the cleaned slice
	for k := range set {
		cleanedSlice = append(cleanedSlice, k)
	}

	return cleanedSlice
}
