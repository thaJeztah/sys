// Package userns provides utilities to detect whether we are currently running
// in a Linux user namespace.

package userns

// RunningInUserNS detects whether we are currently running in a Linux
// user namespace and memoizes the result. It returns false on non-Linux
// platforms.
func RunningInUserNS() bool {
	return inUserNS()
}
