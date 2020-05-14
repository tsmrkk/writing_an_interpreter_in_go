package ast

import "monkey/token"

type Node interface {
	// will be used only for debugging and testing
	TokenLiteral() string
}

type Statement interface {
	Node
	// dummy method
	// helps by guiding the Go compiler and possibly causing it to throw errors
	statementNode()
}

type Expression interface {
	Node
	// dummy method
	// helps by guiding the Go compiler and possibly causing it to throw errors
	expressionNode()
}

// going to be the root node of every AST that the parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token //the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token toke.Token //the token.IDENT token
	Value string
}
