# gen

Gen is a functional programming libaray generator for golang.

### Install

``` go install github.com/whomakemecrazy/gen ```

### Usage 
``` gen -p $packageName -s $serviceName -e $entityName ```

When you run the command, some files are created in the location where you ran the command. (fp.go, $packageName.go)
I usually run it from a domain directory.

EX) user domain generate

``` gen -p user -s service -e User ```


### Directory Tree
```bash
├── app
│   ├── adapter
│   ├── controller
│   ├── mapper
│   ├── presenter
│   └── types
├── command
├── config
│   └── config.go
├── data
├── domain
│   └── user
│     ├── fp.go //generated
│     └── user.go //generated
├── infrastructure
│   ├── database
│   └── network
├── logger
├── protocol
│   └── protocol
├── registry
└── usecase
```
