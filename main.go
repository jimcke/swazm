package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gameplay := Gameplay{}
	newGame(&gameplay)
}
