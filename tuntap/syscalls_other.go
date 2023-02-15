//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package tuntap

import "errors"

func openDev(config Config) (*Interface, error) {
	return nil, errors.New("not implemented on this platform")
}
