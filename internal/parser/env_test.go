package parser_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/ogticrd/kubectl-envsecret/internal/parser"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		setup       func()
		expected    map[string]string
		name        string
		filenames   []string
		expectPanic bool
	}{
		{
			name:      "Single valid file",
			filenames: []string{"testdata/.env1"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/.env1", []byte("KEY1=VALUE1\nKEY2=VALUE2"), 0644)
			},
			expected: map[string]string{
				"KEY1": "VALUE1",
				"KEY2": "VALUE2",
			},
			expectPanic: false,
		},
		{
			name:      "Multiple valid files",
			filenames: []string{"testdata/.env1", "testdata/.env2"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/.env1", []byte("KEY1=VALUE1"), 0644)
				os.WriteFile("testdata/.env2", []byte("KEY2=VALUE2"), 0644)
			},
			expected: map[string]string{
				"KEY1": "VALUE1",
				"KEY2": "VALUE2",
			},
			expectPanic: false,
		},
		{
			name:        "Invalid file name",
			filenames:   []string{"invalid/.env"},
			setup:       func() {},
			expected:    nil,
			expectPanic: true,
		},
		{
			name:      "Mix of valid and invalid files",
			filenames: []string{"testdata/.env1", "invalid/.env"},
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile("testdata/.env1", []byte("KEY1=VALUE1"), 0644)
			},
			expected:    nil,
			expectPanic: true,
		},
		{
			name:        "Empty file list",
			filenames:   []string{},
			setup:       func() {},
			expected:    nil,
			expectPanic: true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			// Setup test environment
			testcase.setup()
			defer os.RemoveAll("testdata")

			// Capture log output
			if testcase.expectPanic {
				log.SetFlags(0)
				log.SetOutput(os.Stderr)
			} else {
				log.SetOutput(os.Stderr)
			}

			defer func() {
				if r := recover(); r != nil {
					if !testcase.expectPanic {
						t.Errorf("unexpected fatal error: %v", r)
					}
				}
			}()

			result, err := parser.Load(testcase.filenames...)
			if !testcase.expectPanic && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if testcase.expectPanic && err == nil {
				t.Errorf("expected fatal error, got nil")
			}
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Errorf("expected %v, got %v", testcase.expected, result)
			}
		})
	}
}
