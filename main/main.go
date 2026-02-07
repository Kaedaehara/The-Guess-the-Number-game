package main

import (
	"bufio"
	"os"
)

func main() {
	g := NewGame()
	g.Run()
}

func NewGame() *Game {
	return &Game{reader: bufio.NewReader(os.Stdin)}
}
