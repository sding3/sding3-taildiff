package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type watcher struct {
	interval  time.Duration
	command   string
	output    *ChangeDetectingBuffer
	exitOnErr bool
	// noNewLine inhibits the printing of a newline character each time
	// the command output changes.
	noNewLine bool
	// withTimestamp causes each output line to be prefixed with the current
	// wall time, which is basically the time that the command responsible for
	// the output line had returned.
	withTimestamp bool
}

func (w *watcher) Start() int {
	tick := setupTicker(w.interval)

	for {
		pState := w.execute()
		if w.exitOnErr && !pState.Success() {
			return pState.ExitCode()
		}

		<-tick
	}
}

func (w *watcher) execute() *os.ProcessState {
	cmd := exec.Command(getShell(), []string{"-c", w.command}...)
	cmd.Stdout = w.output
	cmd.Run() // err discarded
	w.output.Done()

	if w.output.Changed {
		w.print()
	}

	w.output.Rewind()
	return cmd.ProcessState
}

func (w *watcher) print() {
	timestamp := ""
	if w.withTimestamp {
		timestamp = time.Now().Format(time.RFC3339) + " "
	}

	s := w.output.String()
	for l, r := 0, 0; r < len(s); r++ {
		if s[r] == '\n' {
			fmt.Printf("%s%s", timestamp, s[l:r+1])
			l = r + 1
		}
	}

	if !w.noNewLine {
		fmt.Println()
	}
}

func setupTicker(d time.Duration) <-chan time.Time {
	if d <= 0 {
		c := make(chan time.Time)
		close(c)
		return c
	}
	return time.NewTicker(d).C
}

func getShell() string {
	envShell := os.Getenv("SHELL")

	if envShell == "" {
		return "/bin/sh"
	}
	return envShell
}
