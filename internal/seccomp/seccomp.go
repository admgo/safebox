package seccomp

import (
	"os"
	"syscall"

	seccomp "github.com/seccomp/libseccomp-golang"
)

func LoadSeccomp() error {
	err := syscall.Chroot(".")
	if err != nil {
		return err
	}

	err = syscall.Chdir("/")
	if err != nil {
		return err
	}

	// 默认：遇到未允许的 syscall 返回 EPERM（而不是杀进程，便于调试）
	filter, err := seccomp.NewFilter(seccomp.ActKillProcess)
	if err != nil {
		return err
	}
	for _, nr := range ALLOW_SYSCALLS {
		filter.AddRule(seccomp.ScmpSyscall(nr), seccomp.ActAllow)
	}

	// 是否禁用网络（强烈建议禁用）
	if os.Getenv("SECWRAP_NONET") == "1" {
		for _, nr := range ALLOW_NETWORK_SYSCALLS {
			filter.AddRule(seccomp.ScmpSyscall(nr), seccomp.ActErrno.SetReturnCode(1))
		}
	} else {
		for _, nr := range ALLOW_NETWORK_SYSCALLS {
			filter.AddRule(seccomp.ScmpSyscall(nr), seccomp.ActAllow)
		}
	}

	// 强制拒绝的高危调用
	for _, nr := range DENY_SYSCALLS {
		filter.AddRule(seccomp.ScmpSyscall(nr), seccomp.ActErrno.SetReturnCode(1))
	}

	return filter.Load()
}
