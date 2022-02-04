package main

import "conway/app"

func main() {
	app.BenchmarkFromRandom(8192, 0.05, 1000)
	// app.ValidateFromRleFile("_data/random.1024.0.rle", "_data/random.1024.10000.rle", 10000)
}
