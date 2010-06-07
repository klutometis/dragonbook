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

type literal struct {
	rune int
}

func (literal *literal) Value() (interface{}, os.Error) {
	return literal.rune, nil
}

func newLiteral(rune int) *literal {
	return &literal{rune}
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

func readLexeme(rune int,
	size int,
	err os.Error,
	predicate func(rune int) bool,
	reader *bufio.Reader) (string, os.Error) {
	var runes vector.IntVector
	for predicate(rune) && size > 0 && err == nil {
		runes.Push(rune)
		rune, size, err = reader.ReadRune()
	}
	if err != os.EOF {
		unreadRune(reader, rune)
	}
	return lexeme(runes), err
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
			lexeme, err := readLexeme(rune,
				size,
				err,
				unicode.IsDigit,
				reader)
			return newNumber(lexeme), err
		} else if unicode.IsLetter(rune) {
			lexeme, err := readLexeme(rune,
				size,
				err,
				func (rune int) bool { return unicode.IsDigit(rune) ||
					unicode.IsLetter(rune)},
				reader)
			return newSymbol(lexeme), err
		} else {
			return newLiteral(rune), nil
		}
	}
	return nil, os.NewError("this shouldn't happen")
}
