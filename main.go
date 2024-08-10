package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	resetGame()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Крестики нолики"); err != nil {
		log.Fatal(err)
	}
}
