package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/skylark/syntax"
	"github.com/pelletier/go-toml"
)

type dep map[string]string

type Project struct {
	Name     string `toml:"name"`
	Branch   string `toml:"branch,omitempty"`
	Revision string `toml:"revision,omitempty"`
	Version  string `toml:"version,omitempty"`
	Source   string `toml:"source,omitempty"`
}

type Manifest struct {
	Constraints []Project `toml:"constraint,omitempty"`
}

func getDep(s syntax.Stmt) (d dep, ok bool) {
	d, ok = dep{}, false

	e, ok := s.(*syntax.ExprStmt)
	if !ok {
		return
	}
	c, ok := e.X.(*syntax.CallExpr)
	if !ok {
		return
	}
	f, ok := c.Fn.(*syntax.Ident)
	if !ok {
		return
	}
	if f.Name != "go_repository" {
		ok = false
		return
	}

	for _, arg := range c.Args {
		if _, ok := arg.(*syntax.BinaryExpr); !ok {
			continue
		}
		b := arg.(*syntax.BinaryExpr)
		d[b.X.(*syntax.Ident).Name] = b.Y.(*syntax.Literal).Value.(string)
	}
	return d, true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a bazel project path")
		os.Exit(1)
	}

	f, err := syntax.Parse(filepath.Join(os.Args[1], "WORKSPACE"), nil)
	if err != nil {
		panic(err)
	}

	m := Manifest{[]Project{}}
	for _, s := range f.Stmts {
		d, ok := getDep(s)
		if !ok {
			continue
		}

		p := Project{Name: d["importpath"]}
		if commit, ok := d["commit"]; ok {
			p.Revision = commit
		} else if tag, ok := d["tag"]; ok {
			p.Version = tag
		}
		if remote, ok := d["remote"]; ok {
			p.Source = remote
		}
		m.Constraints = append(m.Constraints, p)
	}

	bytes, err := toml.Marshal(m)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filepath.Join(os.Args[1], "Gopkg.toml"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
	file.Close()
}
