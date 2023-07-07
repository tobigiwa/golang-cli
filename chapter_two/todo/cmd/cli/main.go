package main

import (
	todo "cli-go"
	"flag"
	"fmt"
	"os"
)

const todoFilename = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Practising \"Developed for The Pragmatic Bookshelf\"\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	taskFlag := flag.String("task", "", "Task to be includede in the TodoList")
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
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

	case *completeFlag > 0:
		if err := l.Complete(*completeFlag); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *taskFlag != "":
		l.Add(*taskFlag)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
