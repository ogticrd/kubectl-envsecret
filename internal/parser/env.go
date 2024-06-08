// Package parser provides utilities for parsing .env files.
//
// This package includes functions to load environment variables from .env files
// and return them as a map. It uses the godotenv library for reading the .env files.
package parser

import (
	"log"

	"github.com/joho/godotenv"
)

// Load reads the specified .env files and returns their contents as a map.
//
// This function takes one or more filenames as input and reads the environment
// variables from these files. It returns a map where the keys are the variable
// names and the values are the corresponding values from the .env files. If an
// error occurs while reading any of the files, the function logs the error and
// returns the error along with the partially loaded map.
//
// Parameters:
// - filenames: A variadic parameter specifying the .env files to be loaded.
//
// Returns:
// - A map containing the environment variables and their values.
// - An error if any of the files cannot be read.
//
// Example usage:
// envVars, err := parser.Load(".env", ".env.local")
//
//	if err != nil {
//	    fmt.Println("Error loading .env files:", err)
//	}
//
// fmt.Println(envVars)
func Load(filenames ...string) (map[string]string, error) {
	envConfig, err := godotenv.Read(filenames...)
	if err != nil {
		log.Fatal("Error loading file(s):", filenames, "\n Error:", err)
		return envConfig, err
	}

	return envConfig, err
}
