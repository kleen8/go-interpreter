package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    // Identifiers + literals
    IDENT = "IDENT"
    INT = "INT"
    
    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LT = "<"
    GT = ">"

    EQ = "=="
    NOT_EQ = "!="

    // delimiters
    COMMA = ","
    SEMICOLON = ";"
    
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
    TRUE = "TRUE"
    FALSE = "FALSE"
)

var keywords = map[string]TokenType{
    "fn" : FUNCTION,
    "let" : LET,
    "return" : RETURN,
    "if" : IF,
    "else" : ELSE,
    "true" : TRUE,
    "false" : FALSE,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
