/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"github.com/alexpfx/go-pass_manager/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", version, commit, date)
		os.Exit(0)
	}

	cmd.Execute()
}
