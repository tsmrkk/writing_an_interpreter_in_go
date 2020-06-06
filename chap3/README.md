# Chap 3
## 3.1 Giving meaning to symbols

インタープリタの実装によって、`one`が最初に実行され、`two`が次に実行されるか、`two`が最初に実行され、`one`が次に実行されるかは変わる

```
let one = fn() {
  printLine("one");
  return 1;
}

let two = fn() {
  printLine("two");
  return 2;
}

add(one(), two());
```
## 3.2 Strategies to evaluations
Evaluation is where interpreter implementations(regardless of which language they're interpreting)diverge the most.

### The difference between compilers and interpreters
The line between interpreters and compilers is a blurry one. インタープリタは、概念としては実行可能な生成物(機械語とか)を残さないが、コンパイラーは残す.

### A tree-walking interpreter
The archetype of interpreters. Traverse the AST, visit each node and do what the node signifies: print a string, add two numbers, execute function's body - all on the fly.

### Bytecode
Other interpreters also traverse the AST. Instead of interpreting the AST itself they convert it to bytecode. Bytecode is another intermediate representation and a really dense one at that. In general, the opcodes are pretty similar to the mnemonics of most assembly languages(most bytecode definitions contain opcodes for push pop to do stack operation). **Bytecode is not native machine code nor is it assmbly language**. CPUによって直接実行しようと思ってもできない. その代わりに、virtual machineによってinterpretされる. このvirtual machineによるinterpretingもinterpreteの一部である. また、こういったvirtual machineは、本物のマシンやCPU(bytecodeを解釈することができる)をemulateする.

イメージ的には、インタープリタがバイトコードを解釈できるVM(bytecodeを解釈できるOSやマシンの仮想環境)を包含している。

### JIT interpreter/compiler
Instead of executing the operations specified by the bytecode directly in a virtual machine, the virtual machine then compiles the byte code to native machine code, right before its executed

VM上で実行自体を行わずに、VMがbytecodeから、自分のマシンのCPU上にて実行可能なコード(機械語)を生成し、自分のマシン上で実行する方式. TODO: もうちょっと深堀って調べたい

### Cons & Pros
- A tree-walking interpreter is the slowest amongst all approaches, but easy to build and extend
- An interpreter that compiles to bytecode and uses a virtual machine to evaluate said bytecode is going to be a lot faster. But more complicated and harder to build. また、JITコンパイラーであれば、クロスプラットフォームで実行可能な機械語を生成しないといけないというデメリットもある

## 3.3 A tree-walking interpreter

Evaluator sounds mighty and grand, but it will be just one function called "eval". Its job is to evaluate the AST

Psuedocode for "eval"

```
function eval(astNode) {
  if (astNode is integerLiteral) {
    return astNode.integerValue
  } else if (astNode is booleanLiteral) {
    return astNode.booleanValue
  } else if (astNode is infixExpression) {
    leftEvaluated = eval(astNode.Left)
    rightEvaluated = eval(astNode.Right)
    if (astNode.Operator == "+") {
      return leftEvaluated + rightEvaluated
    } else if (astNode.Operator == "-") {
      return leftEvaluated - rightEvaluated
    }
  }
}
```

## 3.4 Representing objects
## 3.5 Evaluating expressions
## 3.6 Conditional
## 3.7 Return statements
## 3.8 Abort! Abort! There's been a mistake! Or: Error handling
## 3.9 Bindings & The environment
## 3.10 Functions & Function calls
## 3.11 Who's taking the trash out
