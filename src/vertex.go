package birds

import (
	"fmt"
	"math"
	"math/rand"
)

type Vertex struct {
	X float64
	Y float64
}

func (vertex Vertex) String() string {
	return fmt.Sprintf("(%v, %v)", vertex.X, vertex.Y)
}

func (vertex Vertex) Distance(otherVertex Vertex) Vector {
	return Vector{
		otherVertex.X - vertex.X,
		otherVertex.Y - vertex.Y,
	}
}

type Vector struct {
	X float64
	Y float64
}

func (vector Vector) String() string {
	return fmt.Sprintf("(%v, %v)", vector.X, vector.Y)
}

func (vector Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vector.X, 2) + math.Pow(vector.Y, 2))
}

func (vector Vector) Multiply(n float64) Vector {
	return Vector{
		vector.X * n,
		vector.Y * n,
	}
}

func (vector Vector) Add(otherVector Vector) Vector {
	return Vector{
		vector.X + otherVector.X,
		vector.Y + otherVector.Y,
	}
}

func (vector Vector) Reverse() Vector {
	return Vector{
		-vector.X,
		-vector.Y,
	}
}

func (vector Vector) Unit() Vector {
	magnitude := vector.Magnitude()
	if magnitude == 0 {
		return Vector{0, 0}
	}
	return Vector{
		vector.X / magnitude,
		vector.Y / magnitude,
	}
}

func Resultant(vectors []Vector) Vector {
	result := Vector{0, 0}
	for _, vector := range vectors {
		result.X += vector.X
		result.Y += vector.Y
	}
	return result
}

func RandomVector(magnitude float64) Vector {
	angle := rand.Float64() * (2 * math.Pi)
	x := math.Sin(angle)
	y := math.Cos(angle)
	unitVector := Vector{x, y}
	return unitVector.Multiply(magnitude)
}

func IsZeroVector(vector Vector) bool {
	return vector.X == 0 && vector.Y == 0
}
