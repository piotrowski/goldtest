package example

import (
	goldtest "goldtesting"
	"testing"
)

func TestCreateTree(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		filename string
	}{
		{"Simple example of JSON", 100, "testdata/node100"},
		{"Simple example of JSON", 500, "testdata/node500"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateTree(tt.size)

			goldtest.AssertJSON(t, got, tt.filename)
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		num      [2]int
		filename string
	}{
		{"Simple example of JSON", [2]int{1, 2}, "testdata/sum12"},
		{"Simple example of JSON", [2]int{1234, 1234}, "testdata/sum12341234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.num[0], tt.num[1])

			goldtest.Assert(t, got, tt.filename)
		})
	}
}
