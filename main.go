package main

import (
	"fmt"
	"strings"
)

func lendAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// totalLength, _ := lendAndUpper("paul")
	// fmt.Println(totalLength)
	repeatMe("abc", "def", "ghi", "jkl")
}
