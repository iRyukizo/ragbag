package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	debug    bool
	rom_path string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Allow debug mode.")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: ragbag [options] [file]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "file is mandatory")
		flag.Usage()
		os.Exit(2)
	}
	if len(flag.Args()) > 1 {
		fmt.Fprintln(os.Stderr, "multiple files is not handled")
		flag.Usage()
		os.Exit(2)
	}

	rom_path = flag.Arg(0)
}

func main() {
	fmt.Println("ragbag", rom_path)
}
