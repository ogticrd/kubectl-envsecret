package parser

import (
	"log"

	"github.com/joho/godotenv"
)

func Load(filenames ...string) (map[string]string, error) {
	envConfig, err := godotenv.Read(filenames...)
	if err != nil {
		log.Fatal("Error loading file(s):", filenames, "\n Error:", err)
		return envConfig, err
	}

	return envConfig, err
}
