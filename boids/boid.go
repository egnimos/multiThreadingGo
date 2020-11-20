package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id int
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{x: rand.Float64() * screenWidth1, y: rand.Float64() * screenHeight1},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id: bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	aceleration := b.calcAcceleration()
	rwlock.Lock()
	b.velocity = b.velocity.Add(aceleration).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rwlock.Unlock()
}

//calculate the acceleration of the neighbouring boid
func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition, avgVelocity, seperation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0

	rwlock.RLock()
	//reading the view map
	//iterate over the view box and get the velocity of other boids
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth1); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight1); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					seperation = seperation.Add(b.position.Subtract(boids[otherBoidId].position).DivideV(dist))
				}
			}
		}
	}
	rwlock.RUnlock()

	accel := Vector2D{x: b.borderBounce(b.position.x, screenWidth1), y: b.borderBounce(b.position.y, screenHeight1)}
	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		avgPosition = avgPosition.DivideV(count)

		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeperation := seperation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeperation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos - viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}