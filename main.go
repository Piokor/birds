package main

import (
	"fmt"

	birds "github.com/Piokor/birds/src"
)

func main() {
	b1 := birds.NewBird(
		10, 10, 0, 0,
	)
	b2 := birds.NewBird(
		10, 10, 0, 0,
	)
	flock := birds.Flock{
		Birds: []*birds.Bird{b1, b2},
	}
	for range 10 {
		fmt.Println(flock)
		flock.Update()
	}
}
