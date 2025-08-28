package python

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

type Capturer struct {
	stdout chan []byte
	stderr chan []byte
	done   chan bool
}

func NewOutputCapturer() *Capturer {
	return &Capturer{
		stdout: make(chan []byte),
		stderr: make(chan []byte),
		done:   make(chan bool),
	}
}

func (s *Capturer) WriteError(data []byte) {
	if s.stderr != nil {
		s.stderr <- data
	}
}

func (s *Capturer) WriteOutput(data []byte) {
	if s.stdout != nil {
		s.stdout <- data
	}
}

func (s *Capturer) CaptureOutput(cmd *exec.Cmd) error {
	// create a pipe for the stdout
	stdout_reader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// create a pipe for the stderr
	stderr_reader, err := cmd.StderrPipe()
	if err != nil {
		stdout_reader.Close()
		return err
	}

	// start the process
	err = cmd.Start()
	if err != nil {
		stdout_reader.Close()
		stderr_reader.Close()
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	written := 0

	// read the output
	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stdout_reader.Read(buf)
			// exit if EOF
			if err != nil {
				if err == io.EOF {
					break
				} else {
					s.WriteError([]byte(fmt.Sprintf("error: %v\n", err)))
					break
				}
			}
			written += n
			s.WriteOutput(buf[:n])
		}
	}()

	// read the error
	go func() {
		buf := make([]byte, 1024)
		defer wg.Done()
		for {
			n, err := stderr_reader.Read(buf)
			// exit if EOF
			if err != nil {
				if err == io.EOF {
					break
				} else {
					s.WriteError([]byte(fmt.Sprintf("error: %v\n", err)))
					break
				}
			}
			s.WriteError(buf[:n])
		}
	}()

	// wait for the process to finish
	go func() {
		// wait for the stdout and stderr to finish
		wg.Wait()

		// wait for the process to finish
		status, err := cmd.Process.Wait()
		if err != nil {
			s.WriteError([]byte(fmt.Sprintf("error: %v\n", err)))
		} else if status.ExitCode() != 0 {
			exit_string := status.String()
			if strings.Contains(exit_string, "bad system call") {
				s.WriteError([]byte("error: operation not permitted\n"))
			} else {
				s.WriteError([]byte(fmt.Sprintf("error: %v\n", exit_string)))
			}
		}
		s.done <- true
	}()

	return nil
}

func (s *Capturer) GetStdout() chan []byte {
	return s.stdout
}

func (s *Capturer) GetStderr() chan []byte {
	return s.stderr
}

func (s *Capturer) GetDone() chan bool {
	return s.done
}
