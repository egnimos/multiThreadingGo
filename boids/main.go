package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"sync"
)

const (
	screenHeight1, screenWidth1 = 360, 640
	boidCount = 500
	viewRadius = 15
	adjRate = 0.015
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount] *Boid
	boidMap[screenWidth1 + 1][screenHeight1 + 1] int
	rwlock = sync.RWMutex{}
)

//type Game struct {}
//
//func (game Game) Update(screen *ebiten.Image) error {
//	if !ebiten.IsDrawingSkipped() {
//		for _, boid := range boids {
//			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
//			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
//			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
//			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
//		}
//	}
//	return nil
//}
func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		}
	}
	return nil
}

//func (game Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
//	//screenHeight1 = outsideHeight
//	//screenWidth1 = outsideWidth
//	return outsideWidth, outsideHeight
//}

func main() {
	//iterate over the boidMap and assign them with value of -1
	for i , row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	//RunGame(Game{})
	if err := ebiten.Run(update, screenWidth1, screenHeight1, 2, "Boids in a box"); err != nil {
		log.Fatalln(err)
	}
}



