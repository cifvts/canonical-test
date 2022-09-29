package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout
var MaxRounds = 3

func shred(path string) int {
	var file, err = os.OpenFile(path, os.O_WRONLY, 0644)

	if err != nil {
		fmt.Fprintf(out, "Error opening the file, %s\n", err)
		return 1
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(out, "Error obtaining file size, %s\n", err)
		return 1
	}

	randomB := make([]byte, fileInfo.Size())
	for round := 0; round < MaxRounds; round++ {
		rand.Read(randomB)
		file.Write(randomB)
		file.Seek(0, 0)

		err = file.Sync()
		if err != nil {
			fmt.Fprintf(out, "Error syncing the file, %s\n", err)
			return 1
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Fprintf(out, "Error closing the file, %s\n", err)
		return 1
	}

	err = os.Remove(path)
	if err != nil {
		fmt.Fprintf(out, "Error removing the file, %s\n", err)
		return 1
	}

	return 0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(out, "Wrong number of parameters")
		os.Exit(1)
	}

	os.Exit(shred(os.Args[1]))
}
