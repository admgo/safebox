package main

import (
	"log"
	"os"
	"syscall"

	"github.com/admgo/safebox/internal/seccomp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("用法：secwrap -- <python_path> <script> [args...]")
	}

	seccomp.LoadSeccomp()

	target := os.Args[1]
	args := os.Args[1:]

	env := os.Environ()

	if err := syscall.Exec(target, args, env); err != nil {
		log.Fatalf("exec %s 失败: %v", target, err)
	}
}
