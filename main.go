package main

import "conway/app"

func main() {
	app.BenchmarkFromRandom(8192, 0.3, 10)
	// app.ValidateFromRleFile("_data/random.1024.0.rle", "_data/random.1024.1000.rle", 1000)
}
