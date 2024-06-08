// Package utils provides utility functions for common operations.
//
// This package contains various utility functions that can be used
// across different parts of an application to perform common tasks
// such as removing duplicates from slices, string manipulations, etc.
package utils

import (
	"os"
)

// RemoveDuplicatedStringE removes duplicate strings from a slice.
//
// This function takes a slice of strings as input and returns a new slice
// containing only the unique strings from the original slice. The order
// of the strings in the resulting slice is not guaranteed to be the same
// as in the input slice.
//
// Parameters:
// - s: A slice of strings from which duplicates need to be removed.
//
// Returns:
// - A new slice of strings containing only unique elements from the input slice.
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

// ValidatePaths checks if the given paths exist and are accessible.
//
// This function takes a slice of file paths as input and checks each one to
// ensure it exists and is accessible. If any path does not exist or is not
// accessible, it returns an error.
//
// Parameters:
// - s: A slice of strings, each representing a file path to validate.
//
// Returns:
// - An error if any of the paths do not exist or are not accessible, otherwise nil.
//
// Example usage:
// paths := []string{"/path/to/file1", "/path/to/file2"}
// err := ValidatePaths(paths)
//
//	if err != nil {
//	    fmt.Println("One or more paths are invalid:", err)
//	} else {
//
//	    fmt.Println("All paths are valid.")
//	}
func ValidatePaths(s []string) error {
	for _, path := range s {
		if _, err := os.Stat(path); err != nil {
			return err
		}
	}
	return nil
}

// MapStringToBytes converts a map of strings to a map of byte slices.
//
// This function takes a map with string keys and string values and converts it
// into a map with the same keys but with values converted to byte slices.
//
// Parameters:
// - m: A map with string keys and string values to be converted.
//
// Returns:
// - A new map with string keys and byte slice values.
//
// Example usage:
// inputMap := map[string]string{"username": "admin", "password": "secret"}
// byteMap := MapStringToBytes(inputMap)
// fmt.Println(byteMap) // Output: map[username:[97 100 109 105 110] password:[115 101 99 114 101 116]]
func MapStringToBytes(m map[string]string) map[string][]byte {
	var convertedMap map[string][]byte = make(map[string][]byte)
	for key, value := range m {
		convertedMap[key] = []byte(value)
	}

	return convertedMap
}
