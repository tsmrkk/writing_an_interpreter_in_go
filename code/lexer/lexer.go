package lexer

type Lexer struct {
	input        string //current position in input (points to current char)
	position     int    // current reading position in input (after current char)
	readPosition int    // current char under examination
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
