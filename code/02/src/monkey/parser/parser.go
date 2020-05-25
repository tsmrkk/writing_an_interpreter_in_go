package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	// curToken points to the current token
	curToken token.Token
	// peekToken points to the next token
	// need it for what to do next if curToken does not give sufficient information
	errors    []string
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peeknToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// add an error to errors when the type of peekToken doesn't match the expectation
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Advances both curToken and peekToken
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// curTokenのタイプがreturnかlet以外だったらnilを返す
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// peekTokenした結果がIDENT(識別子(例えば変数名), Identifier)じゃなかったらnilを返す
	// IDENTだった場合は、expectPeekメソッドの中でadvances both curToken and peekTokenされる
	// MEMO expectPeekメソッドの副作用が嫌(advanceをさせるのをこのメソッド内でやりたくない)
	// MEMO expectPeekメソッドはboolを返すだけの方が読みやすい
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// 変数名をstmt.Name(*ast.Identifier)に入れている
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// peekTokenした結果がASSIGN(=)じゃなかったらnilを返す
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO we're skipping the expressions until we encounter a semicolon
	// 値は現状は一旦無視している
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO we're skipping the expressions until we encounter a semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
