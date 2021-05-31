# esbuild-plugin-globalvar

```go
api.Build(api.BuildOptions{
  Plugins: []api.Plugin{GlobalVarPlugin(map[string]string{
  	"react":     "React",
  	"react-dom": "ReactDOM",
  })},
})
```

src:
```js
import 'react';
import ReactDOM from 'react-dom';
```

dist:
```js
// global-var:react
var require_react = __commonJS({
  "global-var:react"(exports, module) {
    module.exports = window["React"];
  }
});

// global-var:react-dom
var require_react_dom = __commonJS({
  "global-var:react-dom"(exports, module) {
    module.exports = window["ReactDOM"];
  }
});

// test.js
var import_react = __toModule(require_react());
var import_react_dom = __toModule(require_react_dom());
```