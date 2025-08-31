package main

import (
	"github.com/admgo/safebox/internal/seccomp"
)

import "C"

//export LoadSeccomp
func LoadSeccomp() {
	seccomp.LoadSeccomp()
}

func main() {}
