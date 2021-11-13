package main

import (
	"log"

	"github.com/tkeel-io/tkeel-template/cmd/project"
	"github.com/tkeel-io/tkeel-template/cmd/proto"
	"github.com/tkeel-io/tkeel-template/pkg/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tkeel-template",
	Short:   "An elegant toolkit for Go microservices.",
	Long:    `An elegant toolkit for Go microservices.`,
	Version: version.Version,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
