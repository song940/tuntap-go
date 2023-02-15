//go:build darwin && go1.11
// +build darwin,go1.11

package tuntap

import "syscall"

func setNonBlock(fd int) error {
	return syscall.SetNonblock(fd, true)
}
