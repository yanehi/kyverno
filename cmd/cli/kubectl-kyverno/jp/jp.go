package jp

import (
	"github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/jp/function"
	"github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/jp/parse"
	"github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/jp/query"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "jp",
		Short: `Provides a command-line interface to JMESPath, enhanced with Kyverno specific custom functions.
	      For more information visit: https://kyverno.io/docs/writing-policies/jmespath/.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(query.Command())
	cmd.AddCommand(function.Command())
	cmd.AddCommand(parse.Command())
	return cmd
}