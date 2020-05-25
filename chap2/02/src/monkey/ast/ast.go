package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	// will be used only for debugging and testing
	TokenLiteral() string
	String() string
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

// can be added to the Statements slice of ast.Program
type ExpressionStatement struct {
	Token      token.Token //the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type ReturnStatement struct {
	Token       token.Token //the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// going to be the root node of every AST that the parser produces
type Program struct {
	Statements []Statement
}

// for printing AST nodes for debugging and to compare them with other AST nodes
// really handy and useful in writing tests
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		// appends s.String to out(buffer)
		out.WriteString(s.String())
	}
	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string {
	return i.Value
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
