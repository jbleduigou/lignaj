package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		input    []Rann
		expected Hollad
	}{
		{
			name:  "Empty Input",
			input: []Rann{},
			expected: Hollad{
				Tudennou: []Tudenn{},
			},
		},
		{
			name: "Single Rann with Unique Tudenn",
			input: []Rann{
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 10},
					},
				},
			},
			expected: Hollad{
				Tudennou: []Tudenn{
					{Anv: "Alice", Niver: 10},
				},
			},
		},
		{
			name: "Single Rann with Duplicate Tudenn",
			input: []Rann{
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 10},
						{Anv: "Alice", Niver: 15},
					},
				},
			},
			expected: Hollad{
				Tudennou: []Tudenn{
					{Anv: "Alice", Niver: 25},
				},
			},
		},
		{
			name: "Multiple Ranns",
			input: []Rann{
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 10},
						{Anv: "Bob", Niver: 20},
					},
				},
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 5},
						{Anv: "Charlie", Niver: 30},
					},
				},
			},
			expected: Hollad{
				Tudennou: []Tudenn{
					{Anv: "Charlie", Niver: 30},
					{Anv: "Bob", Niver: 20},
					{Anv: "Alice", Niver: 15},
				},
			},
		},
		{
			name: "Mixed Data",
			input: []Rann{
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 10},
						{Anv: "Bob", Niver: 5},
					},
				},
				{
					Tudennou: []Tudenn{
						{Anv: "Alice", Niver: 20},
						{Anv: "Charlie", Niver: 30},
					},
				},
				{
					Tudennou: []Tudenn{
						{Anv: "Bob", Niver: 15},
						{Anv: "Charlie", Niver: 5},
					},
				},
			},
			expected: Hollad{
				Tudennou: []Tudenn{
					{Anv: "Charlie", Niver: 35},
					{Anv: "Alice", Niver: 30},
					{Anv: "Bob", Niver: 20},
				},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.input)
			if len(got.Tudennou) != len(tt.expected.Tudennou) {
				t.Errorf("sum() = %v, want %v", got, tt.expected)
				return
			}
			for i, v := range got.Tudennou {
				if v != tt.expected.Tudennou[i] {
					t.Errorf("sum() = %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}
