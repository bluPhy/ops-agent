// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/GoogleCloudPlatform/ops-agent/confgenerator/filter/internal/token"
)

const (
	NoState    = -1
	NumStates  = 102
	NumSymbols = 105
)

type Lexer struct {
	src     []byte
	pos     int
	line    int
	column  int
	Context token.Context
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:     src,
		pos:     0,
		line:    1,
		column:  1,
		Context: nil,
	}
	return lexer
}

// SourceContext is a simple instance of a token.Context which
// contains the name of the source file.
type SourceContext struct {
	Filepath string
}

func (s *SourceContext) Source() string {
	return s.Filepath
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	lexer := NewLexer(src)
	lexer.Context = &SourceContext{Filepath: fpath}
	return lexer, nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = &token.Token{}
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		tok.Pos.Context = l.Context
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn
	tok.Pos.Context = l.Context

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '.'
1: ':'
2: 'N'
3: 'O'
4: 'T'
5: '('
6: ')'
7: ','
8: '<'
9: '>'
10: '>'
11: '='
12: '<'
13: '='
14: '!'
15: '='
16: '='
17: '~'
18: '!'
19: '~'
20: '='
21: '-'
22: '+'
23: '~'
24: '\'
25: '"'
26: '"'
27: 'O'
28: 'R'
29: 'A'
30: 'N'
31: 'D'
32: '"'
33: ' '
34: '\r'
35: '\t'
36: '\f'
37: \u00a0
38: '\n'
39: '!'
40: '\'
41: 'a'
42: 'b'
43: 'f'
44: 'n'
45: 'r'
46: 't'
47: 'v'
48: '\'
49: '-'
50: '\'
51: 'u'
52: '\'
53: '\'
54: '\'
55: '\'
56: 'x'
57: '!'
58: '*'
59: '/'
60: ';'
61: '?'
62: '@'
63: '['
64: ']'
65: '+'
66: '-'
67: '\'
68: ','
69: '\'
70: ':'
71: '\'
72: '='
73: '\'
74: '<'
75: '\'
76: '>'
77: '\'
78: '+'
79: '\'
80: '~'
81: '\'
82: '"'
83: '\'
84: '\'
85: '\'
86: '.'
87: '\'
88: '*'
89: '#'-'['
90: ']'-'~'
91: \u00a1-\ufffe
92: '0'-'3'
93: '0'-'7'
94: '0'-'7'
95: '0'-'7'
96: '0'-'7'
97: '0'-'7'
98: '0'-'9'
99: 'a'-'f'
100: 'A'-'F'
101: '#'-'''
102: 'A'-'Z'
103: '^'-'}'
104: .
*/