package main

import (
	"log"

	birds "github.com/Piokor/birds/src"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	birdsNum := 100
	birdsList := make([]*birds.Bird, birdsNum)
	for i := range birdsNum {
		birdsList[i] = birds.NewRandomBird(birds.SCREEN_WIDTH, birds.SCREEN_HEIGHT)
	}
	flock := birds.Flock{
		Birds: birdsList,
	}
	ebiten.SetWindowSize(birds.SCREEN_WIDTH, birds.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Birds")
	if err := ebiten.RunGame(birds.NewGame(flock)); err != nil {
		log.Fatal(err)
	}
}
