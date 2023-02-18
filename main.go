package main

import (
	"fmt"
	"os"

	"github.com/deadManAlive/golaf/util"
	"github.com/deadManAlive/golaf/wav"
)

func main() {
	var filename string

	{
		args := os.Args[1:]
		l := len(args)

		if l < 1 {
			fmt.Println("Error: no file provided.")
			return
		} else if l > 1 {
			fmt.Println("Warning: get multiple args, using the first.")
		}

		filename = args[0]
	}

	riff, err := wav.ReadFile(filename)

	util.Check(err)

	fmt.Printf("%+v\n", riff)
}
