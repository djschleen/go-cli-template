// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	//Afs stores a global OS Filesystem that is used throughout hookz
	Afs   = &afero.Afero{Fs: afero.NewOsFs()}
	debug bool
	//Verbose determines if the execution of hookz should output verbose information
	Verbose bool
	//VerboseOutput is set to true if you wish to see debug information in the hooks as they execute in bash
	VerboseOutput bool
	rootCmd       = &cobra.Command{
		Use:     "go-cli-template",
		Short:   `A go CLI template using cobra/spf`,
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				log.SetOutput(ioutil.Discard)
			}
			util.DoIf(Verbose, func() {
				fmt.Println()
				color.Style{color.FgWhite, color.OpBold}.Println("▄▀▀ █   █   ▀█▀ ██▀ █▄ ▄█ █▀▄ █   ▄▀▄ ▀█▀ ██▀")
				color.Style{color.FgWhite, color.OpBold}.Println("▀▄▄ █▄▄ █    █  █▄▄ █ ▀ █ █▀  █▄▄ █▀█  █  █▄▄")
				fmt.Println()
				fmt.Println("A go CLI template using cobra/spf")
				fmt.Println("https://github.com/djschleen")
				fmt.Printf("Version: %s\n", version)
				fmt.Println()
			})
		},
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.PrintErr(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", true, "show verbose output")
}
