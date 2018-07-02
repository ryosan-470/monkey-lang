package token

// TokenType により、トークンタイプの値を識別する
type TokenType string

// Token は Type 属性とトークンの実際のリテラルを保存するフィールド
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL はトークンが未知である
	ILLEGAL = "ILLEGAL"
	// EOFはコードの終わりを示す
	EOF = "EOF"

	// IDENT は 識別子を表す
	IDENT = "IDENT"
	// INT は数字
	INT = "INT"

	// ASSIGN は識別子
	ASSIGN = "="
	// PLUS 演算子
	PLUS = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "function"
	LET      = "LET"
)
