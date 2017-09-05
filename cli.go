package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"strconv"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		w int
		h int

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.IntVar(&w, "w", 100, "")
	flags.IntVar(&h, "h", 100, "")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = w

	_ = h

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	filename := "img" + strconv.Itoa(w) + "x" + strconv.Itoa(h) + ".jpg"
	file, _ := os.Create(filename)
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}
