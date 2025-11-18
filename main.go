package main

import (
	"flag"
	"fmt"
)

func main() {
	dirPtr := flag.String("d", ".", "The directory to use the command in.")
	flag.Parse()
	fmt.Println(*dirPtr)

	crawl_dir(dirPtr)

}
