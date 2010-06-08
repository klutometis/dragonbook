package lexer

import (
	"bufio"
	"container/vector"
	"os"
	"strconv"
	"unicode"
	"utf8"
)

// this isn't going to work; we might have to have
// const (
// 	NUM = new(number)
// )

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type LexerState struct {
	rune int
	line int
}

func NewLexerState(rune int, line int) *LexerState {
	return &LexerState{rune, line}
}

func CopyLexerState(state *LexerState) *LexerState {
	return &LexerState{state.rune, state.line}
}

type token interface {
	Value() (interface{}, os.Error)
}

type tokenFields struct {
	state *LexerState
	lexeme string
}

type number struct {
	tokenFields
}

func (number *number) Value() (interface{}, os.Error) {
	return strconv.Atoi(number.lexeme)
}

func newNumber(state LexerState, lexeme string) *number {
	return &number{tokenFields{&state, lexeme}}
}

type symbol struct {
	tokenFields
}

func (symbol *symbol) Value() (interface{}, os.Error) {
	return symbol.lexeme, nil
}

func newSymbol(state LexerState, lexeme string) *symbol {
	return &symbol{tokenFields{&state, lexeme}}
}

type literal struct {
	tokenFields
}

func (literal *literal) Value() (interface{}, os.Error) {
	return literal.lexeme, nil
}

func newLiteral(state LexerState, lexeme string) *literal {
	return &literal{tokenFields{&state, lexeme}}
}

func unreadBytes(state *LexerState, reader *bufio.Reader, bytes int) {
	for i := 0; i < bytes; i++ {
		state.rune -= 1
		reader.UnreadByte()
	}
}

func unreadRune(state *LexerState, reader *bufio.Reader, rune int) {
	unreadBytes(state, reader, utf8.RuneLen(rune))
}

func lexeme(runes vector.IntVector) string {
	return string(runes.Data())
}

func readLexeme(state *LexerState,
	rune int,
	size int,
	err os.Error,
	predicate func(rune int) bool,
	reader *bufio.Reader) (string, os.Error) {
	var runes vector.IntVector
	for predicate(rune) && size > 0 && err == nil {
		runes.Push(rune)
		rune, size, err = readRune(state, reader)
	}
	if err != os.EOF {
		unreadRune(state, reader, rune)
	}
	return lexeme(runes), err
}

func readRune(state *LexerState,
	reader *bufio.Reader) (rune int, size int, err os.Error) {
	rune, size, err = reader.ReadRune()
	if size > 0 {
		state.rune += 1
	}
	return
}

func Lexan(state *LexerState) (token, os.Error) {
	for {
		rune, size, err := readRune(state, reader)
		if err != nil {
			return nil, err
		} else if size == 0 {
			return nil, os.EOF
		} else if rune == ' ' || rune == '\t' {
		} else if rune == '\n' {
			state.line += 1
			state.rune = 0
		} else if unicode.IsDigit(rune) {
			lexeme, err := readLexeme(state,
				rune,
				size,
				err,
				unicode.IsDigit,
				reader)
			return newNumber(*state, lexeme), err
		} else if unicode.IsLetter(rune) {
			lexeme, err := readLexeme(state,
				rune,
				size,
				err,
				func (rune int) bool { return unicode.IsDigit(rune) ||
					unicode.IsLetter(rune)},
				reader)
			return newSymbol(*state, lexeme), err
		} else {
			return newLiteral(*state, string(rune)), err
		}
	}
	return nil, os.NewError("this shouldn't happen")
}
