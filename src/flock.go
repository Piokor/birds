package birds

import (
	"fmt"
	"strings"
)

type Flock struct {
	Birds []*Bird
}

func (flock Flock) String() string {
	birdLabels := make([]string, len(flock.Birds))
	for i, bird := range flock.Birds {
		birdLabels[i] = bird.String()
	}
	return fmt.Sprintf("Flock:\n %v", strings.Join(birdLabels, "\n "))
}

func (flock Flock) PointOfMass() Vertex {
	xsum := .0
	ysum := .0

	numBirds := float64(len(flock.Birds))
	if numBirds == 0 {
		return Vertex{0, 0}
	}

	for _, bird := range flock.Birds {
		xsum += bird.Position.X
		ysum += bird.Position.Y
	}

	return Vertex{
		xsum / numBirds,
		ysum / numBirds,
	}
}

func (flock Flock) Direction() Vector {
	xsum := 0.
	ysum := 0.

	numBirds := float64(len(flock.Birds))
	if numBirds == 0 {
		return Vector{0, 0}
	}
	for _, bird := range flock.Birds {
		xsum += bird.Velocity.X
		ysum += bird.Velocity.Y
	}
	result := Vector{
		xsum / numBirds,
		ysum / numBirds,
	}
	return result.Unit()
}

func (flock Flock) Update() error {
	for _, bird := range flock.Birds {
		bird.Move()
	}
	pointOfMass := flock.PointOfMass()
	direction := flock.Direction()
	for _, bird := range flock.Birds {
		err := bird.Adjust(flock.Birds, pointOfMass, direction)
		if err != nil {
			return err
		}
	}
	return nil
}
