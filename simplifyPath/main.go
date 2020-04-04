package main

import (
	"strings"
)

func simplifyPath(path string) string {
	directories := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, d := range directories {
		if d == "" || d == "." {
			continue
		} else if d == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		} else {
			stack = append(stack, d)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	input := "/home//foo/"
	simplifyPath(input)
}
