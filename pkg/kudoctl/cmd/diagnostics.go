package cmd

import (
	"fmt"
	"github.com/kudobuilder/kudo/pkg/kudoctl/cmd/diagnostics"
	"github.com/spf13/cobra"
)

const (
	diagCollectExample = `  # collect diagnostics example
  kubectl kudo diagnostics collect --instance=%instance% --namespace=%namespace%
`
	diagAnalyzeExample = ` # analyze diagnostics example
  kubectl kudo diagnostics analyze cassandra_diagnostics.zip
`
)
func newDiagnosticsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "diagnostics",
		Short:   "diagnostics",
		Long:    "diagnostics command has sub-commands to collect and analyze diagnostics data",
	}
	cmd.AddCommand(newDiagnosticsCollectCmd())
	cmd.AddCommand(newDiagnosticsAnalyzeCmd())
	return cmd
}

func newDiagnosticsCollectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "collect",
		Short:   "collect",
		Long:    "collect diagnostics",
		Example: diagCollectExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return diagnostics.Collect(cmd, &Settings)
		},
	}
	cmd.Flags().StringVar(&diagnostics.Options.Instance, "instance", "", "The instance name.")
	return cmd
}

func newDiagnosticsAnalyzeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "analyze",
		Short:   "analyze",
		Long:    "analyze diagnostics data",
		Example: diagAnalyzeExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("analyze diagnostics - not yet implemented")
			return nil
		},
	}
}
