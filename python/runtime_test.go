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

print("hello")
`)

func TestRuntime_Exec(t *testing.T) {
	pyruntime := NewPythonRuntime()
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
			fmt.Printf("%s", string(out))
			stdout_str += string(out)
		case er := <-stderr:
			fmt.Printf("%s", string(er))
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
	pyruntime := NewPythonRuntime()
	genCode, err := pyruntime.dump(codeS)
	fmt.Println(genCode)
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
