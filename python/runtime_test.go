package python

import (
	"fmt"
	"testing"
	"time"
)

func TestRuntime_Exec(t *testing.T) {
	pyruntime := &PythonRuntime{}
	stdout, stderr, done, err := pyruntime.Run("s")
	if err != nil {
		t.Error(err)
	}
	stdout_str := ""
	stderr_str := ""

	defer close(done)
	defer close(stdout)
	defer close(stderr)

	for {
		select {
		case <-done:
			fmt.Sprintf("stdout: %s\nstderr: %s\n", stdout_str, stderr_str)
		case out := <-stdout:
			stdout_str += string(out)
		case err := <-stderr:
			stderr_str += string(err)
		}
	}
}
func TestLongTimeOutput(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(30 * time.Second)
	defer ticker.Stop()
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("long msg")
		case <-timer.C:
			return
		}
	}
}
