//go:build release

package connecttest

import (
	"embed"

	"github.com/shibukawa/frontend-go"
)

//go:embed frontend/dist/*
var asset embed.FS

func init() {
	frontend.SetFrontAsset(asset, frontend.Opt{
		FrameworkType: frontend.SolidJS,
	})
}
