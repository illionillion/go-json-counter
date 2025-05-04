package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementTotal(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		expected int
	}{
		{"starts at 0", 0, 1},
		{"starts at 5", 5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter{Count: tt.initial}
			c.IncrementTotal()
			assert.Equal(t, tt.expected, c.Count)
		})
	}
}

func TestIncrementByName(t *testing.T) {
	tests := []struct {
		name     string
		initial  []NameCount
		arg      string
		expected []NameCount
	}{
		{
			name:    "adds new name",
			initial: []NameCount{},
			arg:     "alice",
			expected: []NameCount{
				{Name: "alice", Count: 1},
			},
		},
		{
			name: "increments existing name",
			initial: []NameCount{
				{Name: "bob", Count: 2},
			},
			arg: "bob",
			expected: []NameCount{
				{Name: "bob", Count: 3},
			},
		},
		{
			name: "does not affect other names",
			initial: []NameCount{
				{Name: "bob", Count: 1},
				{Name: "alice", Count: 2},
			},
			arg: "bob",
			expected: []NameCount{
				{Name: "bob", Count: 2},
				{Name: "alice", Count: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter{Data: tt.initial}
			c.IncrementByName(tt.arg)
			assert.Equal(t, tt.expected, c.Data)
		})
	}
}
