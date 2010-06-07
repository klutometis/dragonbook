package main

import (
	"./lexer"
	"fmt"
)

func main() {
	state := lexer.NewLexerState(0, 0)
	fmt.Print(lexer.Lexan(state))
	fmt.Print(state)
	fmt.Print(lexer.Lexan(state))
	fmt.Print(state)
	fmt.Print(lexer.Lexan(state))
	fmt.Print(state)
	fmt.Print(lexer.Lexan(state))
	fmt.Print(state)
	fmt.Print(lexer.Lexan(state))
	fmt.Print(state)
	fmt.Print(lexer.Lexan(state))
}
