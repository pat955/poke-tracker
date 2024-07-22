package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func title(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		fmt.Println(r)
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}
