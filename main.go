package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	dirPtr := flag.String("d", ".", "The directory to use the command in.")
	flag.Parse()

	start := time.Now()
	crawlDir(dirPtr)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total time elapsed: %v\n", elapsed)

}
