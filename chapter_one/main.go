package main

import (
	chapter0ne "cli-go/cli"
	"flag"
	"fmt"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")

	flag.Parse()

	fmt.Println(chapter0ne.CountFuncChapterOne(os.Stdin, *lines))
}
