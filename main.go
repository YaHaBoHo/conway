package main

import "conway/app"

func main() {
	app.BenchmarkFromRandom(8192, 0.3, 100, 8)
	// app.ValidateFromRleFile("data/random.1024.0.rle", "data/random.1024.10000.rle", 10000, 8)
}
