package main

import (
	ascart "ascart/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Error! Incorrect parametrs count!")
		return
	}
	var style string
	content := os.Args[1]
	if CheckSymbol(content) {
		fmt.Println("Error! Not ASCII symbol")
		return
	}
	if len(os.Args) == 3 {
		style = os.Args[2]
	}
	if len(os.Args[1]) > 10 {
		fmt.Println("Attention! Your input is too longer, symbols may not display correctly!")
	}
	ascart.ReadFile(content, style)
}

func CheckSymbol(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 31 || s[i] > 127 {
			return true
		}
	}
	return false
}
