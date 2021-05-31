package globalvar

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"strings"
)

const NAMESPACE = "global-var"

func GlobalVarPlugin(m map[string]string) api.Plugin {
	return api.Plugin{
		Name: NAMESPACE,
		Setup: func(build api.PluginBuild) {
			var filters []string
			for k := range m {
				filters = append(filters, k)
			}
			build.OnResolve(api.OnResolveOptions{Filter: fmt.Sprintf(`^(%s)$`, strings.Join(filters, "|"))},
				func(args api.OnResolveArgs) (api.OnResolveResult, error) {
					return api.OnResolveResult{
						Path:      args.Path,
						Namespace: NAMESPACE,
					}, nil
				})
			build.OnLoad(api.OnLoadOptions{Filter: ".*", Namespace: NAMESPACE},
				func(args api.OnLoadArgs) (api.OnLoadResult, error) {
					contents := fmt.Sprintf(`
module.exports = window["%s"]
`, m[args.Path])
					return api.OnLoadResult{
						Contents: &contents,
						Loader:   api.LoaderJS,
					}, nil
				})
		},
	}
}
