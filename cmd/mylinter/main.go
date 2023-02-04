package main

import (
	"github.com/YusukeHino/mylinter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mylinter.Analyzer) }
