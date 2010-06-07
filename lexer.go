package lexer

import (
	"./global"
	"bufio"
	// "bytes"
	"container/vector"
	// "fmt"
	// "io"
	// "io/ioutil"
	"os"
	"strconv"
	"unicode"
	"utf8"
)

var lexbuf [global.BSIZE]byte
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type token interface {
	Value() (interface{}, os.Error)
}

type number struct {
	lexeme string
}

func (number *number) Value() (interface{}, os.Error) {
	return strconv.Atoi(number.lexeme)
}

func newNumber(lexeme string) *number {
	return &number{lexeme}
}

type symbol struct {
	lexeme string
}

func (symbol *symbol) Value() (interface{}, os.Error) {
	return symbol.lexeme, nil
}

func newSymbol(lexeme string) *symbol {
	return &symbol{lexeme}
}

func unreadBytes(reader *bufio.Reader, bytes int) {
	for i := 0; i < bytes; i++ {
		reader.UnreadByte()
	}
}

func unreadRune(reader *bufio.Reader, rune int) {
	unreadBytes(reader, utf8.RuneLen(rune))
}

func lexeme(runes vector.IntVector) string {
	return string(runes.Data())
}

func readLexeme(predicate func()) {
}

func Lexan() (token, os.Error) {
	for {
		rune, size, err := reader.ReadRune()
		if err != nil {
			return nil, err
		} else if size == 0 {
			return nil, os.EOF
		} else if rune == ' ' || rune == '\t' {
		} else if rune == '\n' {
			global.Lineno += 1
		} else if unicode.IsDigit(rune) {
			var runes vector.IntVector
			for unicode.IsDigit(rune) && size > 0 {
				runes.Push(rune)
				rune, size, _ = reader.ReadRune()
			}
			unreadRune(reader, rune)
			return newNumber(lexeme(runes)), nil
		} else if unicode.IsLetter(rune) {
			var runes vector.IntVector
			for unicode.IsDigit(rune) || unicode.IsLetter(rune) {
				runes.Push(rune)
				rune, size, _ = reader.ReadRune()
			}
			unreadRune(reader, rune)
			return newSymbol(lexeme(runes)), nil
		}
	}
	return nil, os.NewError("this shouldn't happen")
}
