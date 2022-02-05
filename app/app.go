package app

import (
	"conway/conway"
	"conway/rle"
	"fmt"
	"math"
	"time"
)

func benchmark(world *conway.World, rounds int, concurrency int) {
	start := time.Now()
	conway.Simulate(world, rounds, concurrency)
	taken := time.Now().Sub(start).Seconds()
	fmt.Printf("Size              : %vx%v\n", world.Size, world.Size)
	fmt.Printf("Concurrency       : %v \n", concurrency)
	fmt.Printf("Time              : %v s\n", math.Round(taken*100)/100)
	fmt.Printf("Rounds            : %v\n", rounds)
	fmt.Printf("Round time (avg)  : %v ms\n", int(1000*taken)/rounds)
	fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(float64(world.Cells*rounds)/(taken*10000))/100)
}

func BenchmarkFromRandom(size int, density float32, rounds int, concurrency int) {
	world := conway.RandomWorld(size, density)
	benchmark(world, rounds, concurrency)
}

func BenchmarkFromRleFile(input string, rounds int, concurrency int) {
	world := rle.ToWorld(rle.FromFile(input))
	benchmark(world, rounds, concurrency)
}

func ValidateFromRleFile(input string, check string, rounds int, concurrency int) bool {
	world := rle.ToWorld(rle.FromFile(input))
	conway.Simulate(world, rounds, concurrency)
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
