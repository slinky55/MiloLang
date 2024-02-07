package main

import (
	"os"
	"unicode"

  "github.com/slinky55/Milo/token"
)

type Tokenizer struct {
	file   string
	pos    int
	offset int
	curr   byte

	tokens []token.Token
}

func NewTokenizer(filename string) (*Tokenizer, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		file:   string(bytes),
		pos:    0,
		offset: 0,
		curr:   bytes[0],
	}, nil
}

func isAlpha(b byte) bool {
  return unicode.IsLetter(rune(b))
}

func isNum(b byte) bool {
  return unicode.IsDigit(rune(b))
}

func isWhitespace(b byte) bool {
  return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func (t *Tokenizer) nextChar() {
  t.pos++
  t.offset++
  if t.pos >= len(t.file) {
    return
  }
  t.curr = t.file[t.pos]
}

func (t *Tokenizer) NextToken() *token.Token {
  if t.pos >= len(t.file) {
    return &token.Token{
      Type: token.EOF, 
      Literal: "",
    }
  }

  for isWhitespace(t.curr) {
    t.nextChar() 
    if (t.pos >= len(t.file)) {
      return &token.Token{
        Type: token.EOF,
        Literal: "",
      }
    }
  }

  var nt token.Token

	switch t.curr {
  case '=':
    nt = token.Token{
      Type: token.ASSIGN, 
      Literal: string(t.curr),
    }
    break
	case ';':
		nt = token.Token{
      Type: token.SEMICOLON, 
      Literal: string(t.curr),
    }
    break
	case '(':
		nt = token.Token{
      Type: token.LPAREN, 
      Literal: string(t.curr),
    }
    break
  case ')':
    nt = token.Token{
      Type: token.RPAREN, 
      Literal: string(t.curr),
    }
    break
	case ',':
    nt = token.Token{
      Type: token.COMMA, 
      Literal: string(t.curr),
    }   		
    break
	case '{':
    nt = token.Token{
      Type: token.LBRACE, 
      Literal: string(t.curr),
    }
    break
	case '}':
    nt = token.Token{
      Type: token.RBRACE, 
      Literal: string(t.curr),
    }
    break
	case '+':
    nt = token.Token{
      Type: token.PLUS, 
      Literal: string(t.curr),
    }
    break
  case 0: 
    return &token.Token{
      Type: token.EOF, 
      Literal: string(t.curr),
    }
  default:
    if isNum(t.curr) {
      for isNum(t.file[t.offset]) {
        t.offset++ 
      }

      lit := string(t.file[t.pos:t.offset])
      nt = token.Token{
        Type: token.NUMBER,
        Literal: lit,
      }
      t.pos = t.offset
      t.curr = t.file[t.pos]
      return &nt
    }

    if isAlpha(t.curr)  {
      for isAlpha(t.file[t.offset]) || isNum(t.file[t.offset]) {
        t.offset++
      }

      lit := string(t.file[t.pos:t.offset])
      nt.Literal = lit
      t.pos = t.offset
      t.curr = t.file[t.pos]

      // reserved word check
      switch lit {
      case "let":
        nt.Type = token.LET
        break
      default:
        nt.Type = token.IDENT
      }

      return &nt
    }

    nt = token.Token{
      Type: token.ILLEGAL, 
      Literal: string(t.curr),
    }
	}

  t.nextChar()
  return &nt
}

