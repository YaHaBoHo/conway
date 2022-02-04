package main

import (
	"conway/conway"
	"conway/rle"
	"fmt"
	"math"
	"time"
)

func benchmark(world *conway.World, rounds int) {
	start := time.Now()
	conway.Simulate(world, rounds)
	taken := time.Now().Sub(start).Seconds()
	fmt.Printf("Size              : %vx%v\n", world.Size, world.Size)
	fmt.Printf("Time              : %v s\n", math.Round(taken*100)/100)
	fmt.Printf("Rounds            : %v\n", rounds)
	fmt.Printf("Round time (avg)  : %v ms\n", int(1000*taken)/rounds)
	fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(float64(world.Cells*rounds)/(taken*10000))/100)
}

func benchmarkRandom(size int, density float32, rounds int) {
	world := conway.RandomWorld(size, density)
	benchmark(world, rounds)
}

func validate(input string, check string, rounds int) bool {
	world := rle.ToWorld(rle.FromFile(input))
	conway.Simulate(world, rounds)
	expectedRle := rle.FromFile(check)
	simulatedRle := rle.FromWorld(world)
	if expectedRle.Data == rle.FromWorld(world).Data {
		fmt.Println("OK: Simulation matches expected result.")
		return true
	}
	fmt.Println("ERROR: Simulation differs from expected result.")
	fmt.Printf("[EXPECTED]  %s\n", rle.Summary(expectedRle))
	fmt.Printf("[SIMULATED] %s\n", rle.Summary(simulatedRle))
	return false
}

func main() {
	benchmarkRandom(8192, 0.3, 1000)
	// validate("_data/pentomino.32.0.rle", "_data/pentomino.32.32.rle", 32)
	// validate("_data/random.1024.0.rle", "_data/random.1024.1000.rle", 1000)
	// validate("_data/random.1024.0.rle", "_data/random.1024.10000.rle", 10000)
}
