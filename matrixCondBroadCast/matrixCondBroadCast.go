package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result = [matrixSize][matrixSize]int{}
	rwLock = sync.RWMutex{}
	cond = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

//generate the random matrix
func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

//takes the row and evaluate the given column of matrix to produce the result of multiplication of the matrix....
func matrixCol(row int) {
	rwLock.RLock()
	for {
		waitGroup.Done()
		cond.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}
}

func main() {
	fmt.Println("Working.....")
	start := time.Now()
	waitGroup.Add(matrixSize)
	for row := 0; row < matrixSize; row++ {
		go matrixCol(row)
	}

	for i := 0; i<100; i++ {
		waitGroup.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		waitGroup.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Println("Time taken", elapsed)
}