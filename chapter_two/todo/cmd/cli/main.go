package main

import (
	"bufio"
	todo "cli-go"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const todoFilename = ".todo.json"

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil

}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Practising \"Developed for The Pragmatic Bookshelf\"\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	addkFlag := flag.String("add", "", "Task to be includede in the TodoList")
	listFlag := flag.Bool("list", false, "List all task")
	completeFlag := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &todo.List{}

	switch {
	case *listFlag:
		if err := l.Get(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Print(l)

	case *completeFlag > 0:
		if err := l.Complete(*completeFlag); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *addkFlag != "":
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
