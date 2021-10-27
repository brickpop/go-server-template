package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vocdoni/go-api/config"
	"github.com/vocdoni/go-api/service"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "apiservice",
		Short: "apiservice is a public restful API for Vocdoni, exposing high level abstractions for governance service integrators",
		Long:  `apiservice is a public restful API for Vocdoni, exposing high level abstractions for governance service integrators`,
		Run: func(cmd *cobra.Command, args []string) {
			service.Run()
		},
	}

	cobra.OnInitialize(func() {
		config.Init(rootCmd)
	})

	config.DefineCliFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
