//go:build !embedded_frontend

package web

import "io/fs"

func DistFS() (fs.FS, bool) {
	return nil, false
}
