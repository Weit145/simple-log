package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestRunFormatsBothStreamsAndReturnsExitCode(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--", "sh", "-c", `printf '{"msg":"out"}\nplain\n'; printf '{"msg":"err"}\n' >&2; exit 23`}, &stdout, &stderr)

	if code != 23 {
		t.Errorf("exit code = %d, want 23", code)
	}
	if got := stdout.String(); !strings.Contains(got, "msg: out") || !strings.Contains(got, "plain\n") {
		t.Errorf("stdout = %q", got)
	}
	if got := stderr.String(); !strings.Contains(got, "msg: err") {
		t.Errorf("stderr = %q", got)
	}
}

func TestRunRejectsMissingCommand(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := run([]string{"--"}, &stdout, &stderr); code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
}

func TestSignalHelper(t *testing.T) {
	if os.Getenv("SIMPLE_LOG_TEST_SIGNAL_HELPER") != "1" {
		return
	}
	os.Exit(run([]string{"--", "sh", "-c", `trap "printf 'caught TERM\n'; exit 42" TERM; trap "printf 'caught INT\n'; exit 43" INT; printf 'ready\n'; while :; do sleep 0.05; done`}, os.Stdout, os.Stderr))
}

func TestForwardsSignals(t *testing.T) {
	for _, tc := range []struct {
		name string
		sig  syscall.Signal
		code int
	}{
		{"TERM", syscall.SIGTERM, 42},
		{"INT", syscall.SIGINT, 43},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestSignalHelper$")
			cmd.Env = append(os.Environ(), "SIMPLE_LOG_TEST_SIGNAL_HELPER=1")
			pipe, err := cmd.StdoutPipe()
			if err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			if err := cmd.Start(); err != nil {
				t.Fatal(err)
			}
			reader := bufio.NewReader(pipe)
			line, err := reader.ReadString('\n')
			if err != nil || line != "ready\n" {
				t.Fatalf("ready line = %q, error = %v, stderr = %q", line, err, stderr.String())
			}
			if err := cmd.Process.Signal(tc.sig); err != nil {
				t.Fatal(err)
			}
			remaining, err := io.ReadAll(reader)
			if err != nil {
				t.Fatal(err)
			}
			waitErr := cmd.Wait()
			var exitErr *exec.ExitError
			if !errors.As(waitErr, &exitErr) || exitErr.ExitCode() != tc.code {
				t.Fatalf("exit = %v, want %d; output = %q, stderr = %q", waitErr, tc.code, remaining, stderr.String())
			}
			if !strings.Contains(string(remaining), "caught "+tc.name) {
				t.Errorf("output = %q, want caught %s", remaining, tc.name)
			}
		})
	}
}
