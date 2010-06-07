package global

const (
	BSIZE = 128
	NONE = -1
	EOS = '\000'
	NUM = 256
	DIV = 257
	MOD = 258
	ID = 259
	DONE = 260
)

var Tokenval = NONE
var Lineno = 1
type entry struct {
	lexptr *byte
	token int
}
var Symtable []entry
