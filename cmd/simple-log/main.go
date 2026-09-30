package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Weit145/simple-log/internal/parse"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		if err := parse.Parse(os.Stdin, stdout); err != nil {
			fmt.Fprintln(stderr, "read error:", err)
			return 1
		}
		return 0
	}
	if args[0] != "--" || len(args) == 1 {
		fmt.Fprintln(stderr, "usage: simple-log [-- command [args...]]")
		return 2
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdin = os.Stdin
	childStdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(stderr, "stdout pipe:", err)
		return 1
	}
	childStderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(stderr, "stderr pipe:", err)
		return 1
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)
	if err := cmd.Start(); err != nil {
		fmt.Fprintln(stderr, "start command:", err)
		if errors.Is(err, exec.ErrNotFound) {
			return 127
		}
		return 1
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case sig := <-signals:
				_ = cmd.Process.Signal(sig)
			case <-done:
				return
			}
		}
	}()

	var readers sync.WaitGroup
	readers.Add(2)
	go func() {
		defer readers.Done()
		if err := parse.Parse(childStdout, stdout); err != nil {
			fmt.Fprintln(stderr, "stdout read error:", err)
		}
	}()
	go func() {
		defer readers.Done()
		if err := parse.Parse(childStderr, stderr); err != nil {
			fmt.Fprintln(stderr, "stderr read error:", err)
		}
	}()
	readers.Wait()
	err = cmd.Wait()
	close(done)
	if err == nil {
		return 0
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok && status.Signaled() {
			return 128 + int(status.Signal())
		}
		return exitErr.ExitCode()
	}
	fmt.Fprintln(stderr, "wait command:", err)
	return 1
}
