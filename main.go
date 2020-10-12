// Copyright 2020 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const usage = `Usage:
    isgit [-o OUTPUT] [PATH ...]
 
Options:
    -o, --output OUTPUT         Write the result to the file at path OUTPUT.

isgit iterates over all PATH arguments and writes any PATH that
points to a git repository (containing a '.git' subdirectory) to
OUTPUT.
isgit reads a list of file paths from standard input if no PATH
argument(s) have been provided or when one PATH is '-'.

OUTPUT defaults to standard output.

Example:
   $ find -H -L "$HOME" -type d | isgit
   $ fd -H -L -t d "."  "$HOME" | isgit
`

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }

	var outputFile string
	flag.StringVar(&outputFile, "o", "", "Write the result to this file path")
	flag.StringVar(&outputFile, "output", "", "Write the result to this file path")
	flag.Parse()

	var (
		input  = os.Stdin
		output = os.Stdout
		err    error
	)
	if outputFile != "" {
		output, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalf("Error: failed to open %q: %v", outputFile, err)
		}
		defer output.Close()
	}

	var useStdin bool
	for _, path := range flag.Args() {
		if path == "-" {
			useStdin = true
		} else {
			isgit(output, path)
		}
	}
	if flag.NArg() == 0 || useStdin {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			isgit(output, scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func isgit(output *os.File, path string) {
	file, err := os.Lstat(filepath.Join(path, ".git"))
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error: failed to open %q: %v", filepath.Join(path, ".git"), err)
	}
	if err != nil || !file.IsDir() {
		return
	}

	if _, err := fmt.Fprintln(output, path); err != nil {
		if output == os.Stdout {
			log.Fatalf("Error: failed to write to standard output: %v", err)
		}
		log.Fatalf("Error: failed to write to %q: %v", output.Name(), err)
	}
}
