package main

import (
	"flag"
	"os"
	"text/template"
)

func main() {
	var pkg, srv, entity string
	flag.StringVar(&pkg, "p", "pkg", "pkgName")
	flag.StringVar(&srv, "s", "srv", "serviceName")
	flag.StringVar(&entity, "e", "entity", "entityName")
	flag.Parse()

	values := New(pkg, srv, entity)

	tpl, err := template.New("meta").Parse(code)
	if err != nil {
		panic(err)
	}

	var f *os.File

	f, err = os.Create("fp.go")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(f, values)
	if err != nil {
		panic(err)
	}
	f.Close()
}
