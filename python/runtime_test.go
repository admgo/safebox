package python

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"
)

var code = []byte(`
#!/usr/bin/python3
# -*- coding: utf-8 -*-

# @Time: 2025/8/31 16:17
# @Author: Kenley Wang
# @FileName: test.py

def main():
	import time
	print(time.time())
`)

func TestRuntime_Exec(t *testing.T) {
	w := NewWorkspace("/sandbox")
	pyruntime := NewPythonRuntime(w)
	stdout, stderr, done, err := pyruntime.Run(string(code))
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
			fmt.Printf("stdout: %s\nstderr: %s\n", stdout_str, stderr_str)
			return
		case out := <-stdout:
			stdout_str += string(out)
		case er := <-stderr:
			stderr_str += string(er)
		}
	}
}

var codeS = `
#!/usr/bin/python3
# -*- coding: utf-8 -*-

# @Time: 2025/8/31 16:17
# @Author: Kenley Wang
# @FileName: test.py

print("hello")
`

func TestGenerateCode(t *testing.T) {
	w := NewWorkspace("/sandbox")
	pyruntime := NewPythonRuntime(w)
	genCode, err := pyruntime.dump(codeS)
	fmt.Println(genCode)
	if err != nil {
		t.Error(err)
	}
}

func TestInitializeRuntime(t *testing.T) {
	w := NewWorkspace("/sandbox")
	pyruntime := NewPythonRuntime(w)
	err := pyruntime.initialize()
	if err != nil {
		t.Error(err)
	}
}

func TestLongTimeOutput(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(3 * time.Second)
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

func TestWhereisTmp(t *testing.T) {
	fmt.Println(os.TempDir())
}
