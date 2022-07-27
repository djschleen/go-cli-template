package main

import (
	"os"

	"github.com/djschleen/go-cli-template/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
