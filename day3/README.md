# Day 3
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