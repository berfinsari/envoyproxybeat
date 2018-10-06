package main

import (
	"os"

	"github.com/berfinsari/envoyproxybeat/cmd"

	_ "github.com/berfinsari/envoyproxybeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
