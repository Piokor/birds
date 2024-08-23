package birds

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bird struct {
	Position Vertex
	Velocity Vector
	NextMove Vector
}

func NewBird(posX, posY, velX, velY float64) *Bird {
	bird := Bird{
		Vertex{posX, posY},
		Vector{velX, velY},
		Vector{0, 0},
	}
	return &bird
}

func randomPosition(maxVal float64) float64 {
	centerFactor := 0.5
	value := rand.Float64() * (maxVal * centerFactor)
	return value + (maxVal * (centerFactor / 2))
}

func NewRandomBird(maxX, maxY float64) *Bird {
	bird := Bird{
		Vertex{randomPosition(maxX), randomPosition(maxY)},
		Vector{0, 0},
		Vector{0, 0},
	}
	return &bird
}

func (bird Bird) String() string {
	return fmt.Sprintf("Bird on %v, moving with %v", bird.Position, bird.Velocity)
}

func (bird *Bird) neighborsDistances(birds []*Bird) ([]Vector, error) {
	numOfBirds := len(birds)
	result := make([]Vector, numOfBirds-1)
	result_i := 0
	for _, otherBird := range birds {
		if bird == otherBird {
			continue
		}
		if result_i == numOfBirds-1 {
			return nil, errors.New("given bird not part of the flock")
		}
		result[result_i] = bird.Position.Distance(otherBird.Position)
		result_i += 1
	}
	return result, nil
}

func (bird *Bird) separationVector(birds []*Bird) (Vector, error) {
	neigborDistances, err := bird.neighborsDistances(birds)
	if err != nil {
		return Vector{0, 0}, err
	}
	countedNeighbors := make([]Vector, 0, len(neigborDistances))
	for _, distance := range neigborDistances {
		distanceMagnitude := distance.Magnitude()
		if distanceMagnitude > SEPARATION_RANGE {
			continue
		}
		neigborVector := Vector{distance.X, distance.Y}
		if IsZeroVector(neigborVector) {
			neigborVector = RandomVector(1)
		}
		countedNeighbors = append(countedNeighbors, neigborVector.Unit())
	}
	return Resultant(countedNeighbors).Reverse().Unit(), nil
}

func (bird *Bird) nextMove(flockBirds []*Bird, flockPointOfMass Vertex, flockDirection Vector, target Vertex) (Vector, error) {
	separationVector, err := bird.separationVector(flockBirds)
	if err != nil {
		return Vector{0, 0}, err
	}
	separationVector = separationVector.Multiply(SEPARATION_FACTOR)

	pointOfMassVector := bird.Position.Distance(flockPointOfMass).Unit().Multiply(COHESION_FACTOR)

	targetVector := bird.Position.Distance(target).Unit().Multiply(TARGET_FACTOR)

	flockDirectionVector := flockDirection.Multiply(ALINGMENT_FACTOR)

	return Resultant([]Vector{separationVector, pointOfMassVector, flockDirectionVector, targetVector}), nil
}

func (bird *Bird) Move() {
	bird.Position.X += bird.Velocity.X
	bird.Position.Y += bird.Velocity.Y
}

func (bird *Bird) Adjust(flockBirds []*Bird, flockPointOfMass Vertex, flockDirection Vector, target Vertex) error {
	nextMove, err := bird.nextMove(flockBirds, flockPointOfMass, flockDirection, target)
	if err != nil {
		return err
	}
	bird.NextMove = nextMove
	return nil
}

func (bird *Bird) Turn() {
	bird.Velocity = bird.Velocity.Add(bird.NextMove)
	velocityMagnitude := bird.Velocity.Magnitude()
	if MAX_SPEED != 0 && velocityMagnitude > MAX_SPEED {
		bird.Velocity = bird.Velocity.Unit().Multiply(MAX_SPEED / velocityMagnitude)
	}
}

func (bird *Bird) Draw(screen *ebiten.Image) {
	c := color.RGBA{
		R: uint8(0xff),
		G: uint8(0xff),
		B: uint8(0xff),
		A: 0xff}
	vector.DrawFilledCircle(screen, float32(bird.Position.X), float32(bird.Position.Y), 3, c, false)
}
