package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const (
	versionMajor = 0
	versionMinor = 0
	versionPatch = 2
)

var CommandWatcher = &watcher{output: &ChangeDetectingBuffer{}}

var ProgramName = path.Base(os.Args[0])

func init() {
	flag.ErrHelp = fmt.Errorf("flag error")
	flag.StringVar(&CommandWatcher.command, "c", "", "[required] shell command")
	flag.DurationVar(&CommandWatcher.interval, "n", time.Second, "update interval")
	flag.BoolVar(&CommandWatcher.exitOnErr, "e", false, "exit on command error")
	flag.BoolVar(&CommandWatcher.noNewLine, "no-newline", false, "don't print the extra newline character after each output")
	flag.BoolVar(&CommandWatcher.withTimestamp, "time-stamp", false, "prefix timstamp to each output line")

	origUsage := flag.Usage
	flag.Usage = func() {
		fmt.Printf("%s monitors changes of a supplied shell command's stdout.\n\n", ProgramName)
		origUsage()
	}
}

func main() {
	flag.Parse()

	if strings.ToLower(flag.Arg(0)) == "version" {
		printVersion()
		os.Exit(0)
	}

	if err := validateWatcher(CommandWatcher); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(CommandWatcher.Start())
}

func printVersion() {
	fmt.Printf("taildiff version %d.%d.%d\n", versionMajor, versionMinor, versionPatch)
}

func validateWatcher(w *watcher) error {
	if w.command == "" {
		return fmt.Errorf("A command is required. Supply command with -c <command>")
	}
	return nil
}
