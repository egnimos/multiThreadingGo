package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

//add two vectors
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

//subtract two vectors
func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
	}
}

//multiply two vectors
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
	}
}

//add value to a vector
func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{
		x: v1.x + d,
		y: v1.y + d,
	}
}

//multiply value to a vector
func (v1 Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{
		x: v1.x * d,
		y: v1.y * d,
	}
}

//divide a value to a vector
func (v1 Vector2D) DivideV(d float64) Vector2D {
	return Vector2D{
		x: v1.x / d,
		y: v1.y / d,
	}
}

//limit the vector to certain range
func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{
		x: math.Min(math.Max(v1.x, lower), upper),
		y: math.Min(math.Max(v1.y, lower), upper),
	}
}

//distance
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x - v2.x, 2) + math.Pow(v1.y - v2.y, 2))
}


