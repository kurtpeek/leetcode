package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	var testCases = []struct {
		input, output string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			assert.Exactly(t, tc.output, simplifyPath(tc.input))
		})
	}
}
