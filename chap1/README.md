# Chap 1
## Interpreter
Interpreters take source code and evaluate it without producing some visible, intermediate result that can later be executed. That's in contrast to compilers, which take source code and produce output in another language that the underlying system can understand.

## Bytecode
Some interpreters compile input into an internal representation called bytecode

## JIT interpreters
JIT interpreters compile the input just-in-time into native machine code that which gets then executed

## tree-walking interpreters
tree-walking interpreters are interpreters that parse the source code, build an abstract syntax tree(AST) out of it and the evaluate this tree.

## first class functions
Functions are just values like integers or strings in the idea of first class functions

## 1.1 Lexical Analysis
<img src="./images/1.png">

Lexer: Does lexical analysis. It is also called tokernizer or scanner

Lexical Analysis: The first transfomation, from source code to tokens. This process is also called "lexing".

Tokens: Small, easily categorized data structure that are then fed to parser

Parser: Does the second transfomation and turns the tokens in "Abstract Syntax Tree"

Example: Input into Lexer

```
"let x = 5 + 5"

```

Example: Output from Lexer

```
[
  LET,
  IDENTIFIER("x"),
  EQUAL_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```

<!--
TODO Dig down for more information. Add more sentences
-->
What constitutes  a token varies between different lexer implementations. Some lexers, parse "5" to an integer in the parsing stage, while other lexers does not do that when constructing tokens

## 1.2 Defining Our Tokens

Example input into Monkey language lexer

```
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
```

### Token data structure

```
// token/token.go
package token
type TokenType string
type Token struct {
  Type TokenType
  Literal string
}
```

### Token Types

```
// token/token.go
const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

//identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y,...
  INT = "INT" //123456

//operators
  ASSIGN = "="
  PLUS = "+"

//delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

//keywords
  FUNCTION = "FUNCTION"
  LET = "LET"


    )

```

## 1.3 The lexer
Goal of this section is to wirite my own lexer

- source code as input
- output the tokens that represent the source code

- NextToken() will output the next token
- initialize the lexer with source code and then repeatedly call NextToken()

In production, it is better to attach filenames and line numbers to tokens to track down lexing and parsing errors. However, in this book, there will not be any code to do that. Beacuase it is really complicated.

### `position` and `readPosition`
- Both will be used to access characters in input by using them as an index
- The reason for these two "poiters" pointing into input string is that it will be necessary to be able to "peek" further into the input to look after the current character to see what comes up next
- `readPosition` always points to the "next" character in the input
- `position` points to the character in the input that coresspond to the ch byte

## 1.4 Extending our token set and lexer
In this section support for `==`, `!`, `!=`, `-`, `/`, `*`, `<`, `>`, `true`, `false`, `if`, `else`, `return` will be added

## 1.5 Start of a repl
REPL stands for "Read Eval Print Loop". Sometimes the REPL is called "console", sometimes "interactive mode". REPL reads input, sends it to the interpreter for evaluation, prints the result(output) of the interpreter and starts again. Read, Eval, Print, Loop
