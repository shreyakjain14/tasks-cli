package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: taskcli <add|delete|done|list> [args]")
		os.Exit(1)
	}

	tl, err := Load()

	if err != nil {
		fmt.Println("error loading tasks:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: taskcli add <title>")
			os.Exit(1)
		}
		tl.Add(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage: taskcli delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			os.Exit(1)
		}
		tl.Delete(id)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("usage: taskcli delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			os.Exit(1)
		}
		tl.Done(id)
	case "list":
		tl.List()
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}
}
