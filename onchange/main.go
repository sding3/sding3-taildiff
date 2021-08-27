package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := run(ctx, os.Args)
	if errors.Is(err, context.Canceled) {
		fmt.Fprintln(os.Stderr, "\ncancelled")
		return
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return
}

func run(ctx context.Context, args []string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	var recursive bool

	switch len(args) {
	case 1:
		err = watcher.Add(".")
	default:
		err = add(watcher, args[1])
		if strings.HasSuffix(args[1], "...") {
			recursive = true
		}
	}
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err = <-watcher.Errors:
				fmt.Fprintln(os.Stderr, err)
			case <-ctx.Done():
				return
			}
		}
	}()

LOOP:
	for {
		select {
		case e := <-watcher.Events:
			if len(os.Args) > 2 {
				cmd := exec.Command(os.Args[2], os.Args[3:]...)
				cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
				err := cmd.Run()
				if err != nil {
					return err
				}
			} else {
				fmt.Println(e)
			}

			if !recursive {
				continue
			}
			switch e.Op {
			case fsnotify.Create:
				err := watcher.Add(e.Name) // err discarded
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			case fsnotify.Remove:
				watcher.Remove(e.Name) // err discarded
			}
		case <-ctx.Done():
			break LOOP
		}
	}

	return ctx.Err()
}

func add(w *fsnotify.Watcher, path string) error {
	if path == "." {
		return w.Add(path)
	}

	addRecursively := func(p string) error {
		return filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return w.Add(path)
		})
	}

	if path == "..." {
		return addRecursively(".")
	}

	if strings.HasSuffix(path, "...") {
		return addRecursively(strings.TrimSuffix(path, "..."))
	}

	return w.Add(path)
}
