//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package tuntap

// PlatformSpeficParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
