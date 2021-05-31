package globalvar

import (
	"github.com/evanw/esbuild/pkg/api"
	"testing"
)

func TestGlobalVarPlugin(t *testing.T) {
	api.Build(api.BuildOptions{
		Bundle:      true,
		Write:       true,
		EntryPoints: []string{"./test.js"},
		Outfile:     "out.js",
		Plugins: []api.Plugin{GlobalVarPlugin(map[string]string{
			"react":     "React",
			"react-dom": "ReactDOM",
		})},
		Format: api.FormatUMD,
	})
}
