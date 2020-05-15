# Day 5
## 2.4 Parser's first steps: parsing let statements

### peekToken and curToken
They act like the two "pointers" that the lexer we've made has: `position` and `peekPosition`.

### ParseProgram
Construct the root node of the AST, an `*ast.Program`. It then iterates over every token in the input until it encounters an `token.EOF` token. It does this by repeatedly calling `nextToken`. which advances both `p.curToken` and `p.peekToken`
