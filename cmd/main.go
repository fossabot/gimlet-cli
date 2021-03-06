package main

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/gimlet-io/gimlet-cli/commands/chart"
	"github.com/gimlet-io/gimlet-cli/commands/gitops"
	"github.com/gimlet-io/gimlet-cli/commands/manifest"
	"github.com/gimlet-io/gimlet-cli/commands/seal"
	"github.com/gimlet-io/gimlet-cli/version"
	"github.com/urfave/cli/v2"
	"os"
)

//go:generate go run ../scripts/includeWeb.go

func main() {
	app := &cli.App{
		Name:                 "gimlet",
		Version:              version.String(),
		Usage:                "for an open-source GitOps workflow",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			&chart.Command,
			&gitops.Command,
			&seal.Command,
			&manifest.Command,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", emoji.CrossMark, err.Error())
		os.Exit(1)
	}
}
