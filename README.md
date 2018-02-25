This package converts `go_repository` rules in a valid Bazel WORKSPACE file into [dep](https://golang.github.io/dep/) constraints. The `[[constraint]]` rules are appended to the current `Gopkg.toml`, if any.


Usage:
```
go get github.com/crazytan/bazel-to-go
go run $GOPATH/github.com/crazytan/bazel-to-go/main.go <project path>
```


Some implicit dependencies are added automatically if you already use `go_rules_dependencies()` and you may want to add them to `Gopkg.toml` manually. See [here](https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id4) for a complete list.
