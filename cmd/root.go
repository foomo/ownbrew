package cmd

import (
	"os"

	pkgcmd "github.com/foomo/ownbrew/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var root *cobra.Command

func init() {
	root = NewRoot()
	NewConfig(root)
	NewInit(root)
	NewInstall(root)
	NewVersion(root)
	cobra.OnInitialize(pkgcmd.InitConfig)
}

// NewRoot represents the base command when called without any subcommands
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ownbrew",
		Short: "Your local project package manager",
	}
	cmd.PersistentFlags().BoolP("verbose", "v", false, "output debug information")
	_ = viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

	cmd.PersistentFlags().StringP("config", "c", ".ownbrew.yaml", "config file (default is .ownbrew.yaml)")
	_ = viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1) //nolint:revive
	}
}
