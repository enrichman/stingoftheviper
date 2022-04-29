package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewStingCommand(config *Config) *cobra.Command {
	// get the address of your sub-command config, or you will not be able to rewrite its values
	stingConfig := &config.StingConfig

	// Define our command
	stingCmd := &cobra.Command{
		Use:   "sting",
		Short: "a small subcommand",

		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()

			fmt.Fprintln(out, "Verbosity is set to:", config.Verbose)
			fmt.Fprintln(out, "You want to name:", config.Name)
			fmt.Fprintln(out, "You want to sting:", config.StingConfig.Name)
		},
	}

	// bind the root flags and persistenceFlags to the config
	bindStingFlags(stingCmd.Flags(), stingCmd.PersistentFlags(), stingConfig)

	return stingCmd
}

func bindStingFlags(flags, persistentFlags *pflag.FlagSet, config *StingConfig) {
	flags.StringVarP(&config.Name, "name", "n", config.Name, "Who do you want to sting?")
}
