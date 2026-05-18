//go:build embedded_frontend

package web

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var embeddedDist embed.FS

func DistFS() (fs.FS, bool) {
	dist, err := fs.Sub(embeddedDist, "dist")
	if err != nil {
		return nil, false
	}
	return dist, true
}
