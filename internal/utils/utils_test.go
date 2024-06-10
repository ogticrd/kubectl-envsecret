package utils

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestRemoveDuplicatedStringE(t *testing.T) {
	testcases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "No duplicates",
			input:    []string{"red", "green", "blue"},
			expected: []string{"red", "green", "blue"},
		},
		{
			name:     "All duplicates",
			input:    []string{"red", "red", "red"},
			expected: []string{"red"},
		},
		{
			name:     "Some duplicates",
			input:    []string{"red", "green", "red", "blue", "green"},
			expected: []string{"red", "green", "blue"},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"red"},
			expected: []string{"red"},
		},
		{
			name:     "Multiple unique elements",
			input:    []string{"red", "green", "blue", "yellow", "purple"},
			expected: []string{"red", "green", "blue", "yellow", "purple"},
		},
		{
			name:     "Case sensitivity",
			input:    []string{"Red", "red", "Green", "green"},
			expected: []string{"Red", "red", "Green", "green"},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			testResult := RemoveDuplicatedStringE(test.input)
			if !reflect.DeepEqual(testResult, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, testResult)
			}
		})
	}
}

func TestValidatePaths(t *testing.T) {
	tests := []struct {
		expectedErr error
		setup       func()
		name        string
		input       []string
	}{
		{
			name:  "All valid paths",
			input: []string{"testdata/file1.txt", "testdata/file2.txt"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/file1.txt", []byte("content"), 0644)
				os.WriteFile("testdata/file2.txt", []byte("content"), 0644)
			},
			expectedErr: nil,
		},
		{
			name:        "All invalid paths",
			input:       []string{"invalid/path1.txt", "invalid/path2.txt"},
			setup:       func() {},
			expectedErr: os.ErrNotExist,
		},
		{
			name:  "Mix of valid and invalid paths",
			input: []string{"testdata/file1.txt", "invalid/path1.txt"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/file1.txt", []byte("content"), 0644)
			},
			expectedErr: os.ErrNotExist,
		},
		{
			name:        "Empty slice",
			input:       []string{},
			setup:       func() {},
			expectedErr: nil,
		},
		{
			name:  "Single valid path",
			input: []string{"testdata/file1.txt"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/file1.txt", []byte("content"), 0644)
			},
			expectedErr: nil,
		},
		{
			name:        "Single invalid path",
			input:       []string{"invalid/path1.txt"},
			setup:       func() {},
			expectedErr: os.ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test environment
			tt.setup()
			defer os.RemoveAll("testdata")

			err := ValidatePaths(tt.input)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestMapStringToBytes(t *testing.T) {
	tests := []struct {
		expected map[string][]byte
		input    map[string]string
		name     string
	}{
		{
			name: "Multiple key-value pairs",
			input: map[string]string{
				"one":   "1",
				"two":   "2",
				"three": "3",
			},
			expected: map[string][]byte{
				"one":   []byte("1"),
				"two":   []byte("2"),
				"three": []byte("3"),
			},
		},
		{
			name:     "Empty map",
			input:    map[string]string{},
			expected: map[string][]byte{},
		},
		{
			name: "Special characters",
			input: map[string]string{
				"hello":   "world",
				"special": "!@#$%^&*()",
			},
			expected: map[string][]byte{
				"hello":   []byte("world"),
				"special": []byte("!@#$%^&*()"),
			},
		},
		{
			name: "Unicode characters",
			input: map[string]string{
				"unicode": "こんにちは",
			},
			expected: map[string][]byte{
				"unicode": []byte("こんにちは"),
			},
		},
		{
			name: "Empty string values",
			input: map[string]string{
				"empty": "",
			},
			expected: map[string][]byte{
				"empty": []byte(""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapStringToBytes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
