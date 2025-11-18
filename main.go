package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	dirPtr := flag.String("d", ".", "The directory to use the command in.")
	limitPtr := flag.Int("l", 1, "The maximum amount of threads to be used.")
	flag.Parse()

	start := time.Now()
	crawlDir(dirPtr, *limitPtr)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total time elapsed: %v\n", elapsed)

}
