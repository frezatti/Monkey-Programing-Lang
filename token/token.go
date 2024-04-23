package token

type TokenType string

const(
ILLEGAL="ILLEGAL"
EOF = "EOF"

IDENT = "IDENT"
INT = "INT"

//Operators
ASSIGN = "="
PLUS = "+"

//Delimiters
COMMA = ","
SEMICOLON = ";"

LPAREN = "("
RPAREN = ")"
LBRACE = "{"
RBRACE = "}"

//Keywords
FUNCTION = "FUNCTION"
LET = "LET"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
}
func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok;
    }
    return IDENT;
}

type Token struct{
    Type   TokenType

    Literal string
}