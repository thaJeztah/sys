package user

import (
	"os"
)

// This is currently a wrapper around [os.MkdirAll] since currently
// permissions aren't set through this path, the identity isn't utilized.
// Ownership is handled elsewhere, but in the future could be support here
// too.
func mkdirAs(path string, _ os.FileMode, _, _ int, _ bool, _ ...MkdirOpt) error {
	return os.MkdirAll(path, 0)
}
