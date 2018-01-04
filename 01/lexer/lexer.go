package lexer

type Lexer struct {
	input        string
	position     int  // current position in the input (current char)
	readPosition int  // current reading position in the input (after current char)
	ch           byte //current char under examination, ASCII
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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
