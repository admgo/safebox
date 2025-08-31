package main

import (
	"github.com/admgo/safebox/internal/seccomp"
)

func LoadSeccomp() {
	seccomp.LoadSeccomp()
}

func main() {}
