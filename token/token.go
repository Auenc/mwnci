package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Deliminators
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "ffwythiant"
	LET      = "gosod"
	RETURN   = "adfer"
	TRUE     = "gwir"
	FALSE    = "ffug"
	IF       = "os"
	ELSE     = "arall"
)

var keywords = map[string]TokenType{
	"ffwythiant": FUNCTION,
	"gosod":      LET,
	"adfer":      RETURN,
	"gwir":       TRUE,
	"ffug":       FALSE,
	"os":         IF,
	"arall":      ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
