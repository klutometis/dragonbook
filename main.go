package main

import (
	"dragonbook/lexer"
	"fmt"
)

func main() {
	state := lexer.NewLexerState(0, 0)
	fmt.Println(lexer.Lexan(state))
	fmt.Println(state)
	fmt.Println(lexer.Lexan(state))
	fmt.Println(state)
	fmt.Println(lexer.Lexan(state))
	fmt.Println(state)
	fmt.Println(lexer.Lexan(state))
	fmt.Println(state)
	fmt.Println(lexer.Lexan(state))
	fmt.Println(state)
	fmt.Println(lexer.Lexan(state))
}
