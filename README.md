This package helps to migrate a Bazel-managed Go project to a [dep](https://golang.github.io/dep/)-managed one. It converts `go_repository` rules in a valid Bazel WORKSPACE file into `dep` constraints. The `[[constraint]]` rules are appended to the current `Gopkg.toml` file, if any. Some implicit dependencies are added automatically by Bazel if you already use `go_rules_dependencies()` and you may want to add them to `Gopkg.toml` manually. See [here](https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id4) for a complete list.


Usage:
```bash
$ go get github.com/crazytan/bazel-to-go
$ go run $GOPATH/github.com/crazytan/bazel-to-go/main.go <project path>
```

## Example
`WORKSPACE` file:
```
go_repository(
    name = "com_github_spf13_pflag",
    commit = "4c012f6dcd9546820e378d0bdda4d8fc772cdfea",
    importpath = "github.com/spf13/pflag",
)
go_repository(
    name = "com_github_spf13_cobra",
    commit = "f91529fc609202eededff4de2dc0ba2f662240a3",
    importpath = "github.com/spf13/cobra",
)
```
will be converted to 
```toml
[[constraint]]
  name = "github.com/spf13/pflag"
  revision = "4c012f6dcd9546820e378d0bdda4d8fc772cdfea"

[[constraint]]
  name = "github.com/spf13/cobra"
  revision = "f91529fc609202eededff4de2dc0ba2f662240a3"
```
