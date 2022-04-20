package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Build the cobra command that handles our command line tool.
func NewRootCommand() *cobra.Command {
	// initialize our empty/default configuration
	config := NewConfig()

	// Define our root command
	rootCmd := &cobra.Command{
		Use:   "stingoftheviper",
		Short: "Cobra and Viper together at last",
		Long:  `Demonstrate how to get cobra flags to bind to viper properly`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
			out := cmd.OutOrStdout()

			// Print the final resolved value from binding cobra flags and viper config
			fmt.Fprintln(out, "Your favorite number is:", config.Number)
			fmt.Fprintln(out, "Your name is:", config.Name)
		},
	}

	// bind the root flags and persistenceFlags to the config
	bindRootFlags(rootCmd.Flags(), rootCmd.PersistentFlags(), &config)

	rootCmd.AddCommand(NewStingCommand(&config))

	return rootCmd
}

func bindRootFlags(flags, persistentFlags *pflag.FlagSet, config *Config) {
	flags.StringVarP(&config.Name, "name", "n", config.Name, "What's your name?")

	// this flag will be persisted trough the sub-commands
	persistentFlags.IntVarP(&config.Number, "number", "c", config.Number, "Which is your favorite number?")
}
