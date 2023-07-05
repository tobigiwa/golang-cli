package chpone

import (
	"bufio"
	"io"
	
)

func CountFuncChapterOne(r io.Reader, countLines bool) int {

	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	
	for scanner.Scan() {
		wc++
	}
	return wc
}
