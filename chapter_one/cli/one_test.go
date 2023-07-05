package chpone_test

import (
	"bytes"
	chapter0ne "cli-go/cli"

	"testing"
)

func TestCountFunc(t *testing.T) {
	b := bytes.NewBufferString("word1 \n\nword2 word3")

	exp := 3

	res := chapter0ne.CountFuncChapterOne(b, true)

	if res != exp {
		t.Errorf("expect %v, got %v", exp, res)
	}
}
