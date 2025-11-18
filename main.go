package main

import (
	"flag"
)

func main() {
	dirPtr := flag.String("d", ".", "The directory to use the command in.")
	flag.Parse()

	crawlDir(dirPtr)

}
