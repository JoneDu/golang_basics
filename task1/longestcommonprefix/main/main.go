package main

import (
	"fmt"
	"strings"
)

func main() {
	commonPrefix := longestCommonPrefix([]string{"fl00", "flower", "flw"})
	fmt.Printf("commonPrefix: %+v\n", commonPrefix)
}

func longestCommonPrefix(strs []string) string {
	var longestCommonPrefix string
	var prefix string
	var prefixs []string
	var firsStr = strs[0]
	runes := []rune(firsStr)
	for i := range runes {
		prefix += string(runes[i])
		isCommonPrefix := true
		for i := 1; i < len(strs); i++ {
			str := strs[i]
			if !strings.HasPrefix(str, prefix) {
				isCommonPrefix = false
				break
			}
		}
		if isCommonPrefix {
			prefixs = append(prefixs, prefix)
		}
	}
	for _, p := range prefixs {
		if len(p) > len(longestCommonPrefix) {
			longestCommonPrefix = p
		}
	}
	return longestCommonPrefix
}
