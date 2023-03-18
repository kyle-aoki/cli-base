package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	defer MainRecover()
	ProgramConfigure()
	log.Printf("program started with flags: \n%s", ToJsonStr(Flags))

	panic("test panic")
}

func Must[T any](t T, err error) T {
	Check(err)
	return t
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

var ToJsonStr = func(v any) string { return string(Must(json.MarshalIndent(v, "", "  "))) }

type ProgramFlags struct {
	Verbose *bool
	Debug   *bool
}

var Flags = &ProgramFlags{}

func ProgramConfigure() {
	Help := flag.Bool("h", false, "help")
	Flags.Verbose = flag.Bool("v", false, "verbose")
	Flags.Debug = flag.Bool("d", false, "debug")
	flag.Parse()

	if *Help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !*Flags.Verbose {
		log.SetOutput(io.Discard)
	}
}

func MainRecover() {
	if *Flags.Debug {
		return
	}
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
