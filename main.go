package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	downloadPath := filepath.Join(homeDir, "Downloads")
	dirPtr := flag.String("d", downloadPath, "The directory to use the command in.")
	limitPtr := flag.Int("l", 2, "The maximum amount of threads to be used.")
	flag.Parse()

	start := time.Now()
	crawlDir(dirPtr, *limitPtr)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total time elapsed: %v\n", elapsed)

}
