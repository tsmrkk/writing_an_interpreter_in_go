# Day 6
## 式(expression)と文(statement)の違い
式は値を生成し、変数に代入可能なものを指す. 例として以下のようなものがあげられる
- `42`のようなリテラル(ソースコード内に直接値を表記したものを指す)
- `foo`といった変数
- `hoge()`といった関数呼び出し

文は処理する1ステップを指す. 例としてはif文やfor文. 文の処理の一部として式を含むことがある. 以下の例では、isTrueという式が文の中に出てきている

```
const isTrue = true;
if (isTrue) {
}
```

if文は式ではなく、文であるため代入ができない

```
const ff = if (true) {...}
```

jsにはfunction式とfunction文がある

```
function learn() {
  ...
}

const a = function() {
  ...
};


```
## 2.6 Parsing expressions
Everything besides `let` and `return` statements is an expression

## 参考・引用
- [JavaScript Primer 文と式](https://jsprimer.net/basic/statement-expression/)

