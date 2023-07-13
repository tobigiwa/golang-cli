#  A Go tour of CLI

My tour of the book "Powerful Commandline Application in Go" by Ricardo Gerardi, edited by Brian P. Hogan. I tried writing every line of the code from the book while adding a little twist of mine.

I strongly advise this book for anyone starting Go, as a second book ([Learnig Go](https://learning-go-book.dev/) as a first) and a first and probably the last for CLIs in Go.

```go

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Practising \"Developed for The Pragmatic Bookshelf\"\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	addkFlag := flag.String("add", "", "Task to be includede in the TodoList")
	listFlag := flag.Bool("list", false, "List all task")
	completeFlag := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()
```

![](/assets/Screenshot%20from%202023-07-13%2014-02-46.png)