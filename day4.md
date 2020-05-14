# Day 4
## 2.1 Parsers
Parsers take source code as input and produce a data structure which represents the source code. Inputs to parser are the tokens.

### memo
Although each implementations of AST are all pretty similar, there is not one true, universal AST format that's used by every parser

## 2.2 Why not a parser generator?
### Parser generator
Parser genrerators are tools that, when fed with a formal description of a language, produce parser as their output. The output is code that can then be compiled/interprede and itself fed with source code as input to produce a syntax tree. The majority of them use a context-free grammar as their input.

Parsers are exceptionally well suited to being automatically generated. So, it is recommended to use parser generator for production-environment, where correctness and robustness are priorities.

### Context-Free Grammar
A CFG is a aset of rules that describe how to form correct sentences in a language. Common notationl formats of CFGs are the BNF or the EBNF. EBNF stands for(Extended BNF).

## 2.3 Writing a parser for the monkey programming language
There are two main strategies to parse a programming language:

- top-down parsing
- bottom-up parsing

In this section, we will write a recursive descent parser. It's a "top-down operator precedence" parser

The difference between top-down and bottom-up parsers is that the former starts with constructing root node of the AST and then descends while the latter does it the other way around.
