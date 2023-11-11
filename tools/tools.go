package tools

import (
	_ "github.com/cweill/gotests/gotests"       // tests generator
	_ "github.com/fatih/gomodifytags"           // stuct tags utility
	_ "github.com/go-delve/delve/cmd/dlv"       // debugger
	_ "github.com/haya14busa/goplay/cmd/goplay" // playground
	_ "github.com/josharian/impl"               // method stub generator
	_ "golang.org/x/tools/gopls"                // language server
	_ "honnef.co/go/tools/cmd/staticcheck"      // code checker
)
