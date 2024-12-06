package cmd

import (
	pkgcmd "github.com/foomo/ownbrew/pkg/cmd"
	"github.com/foomo/ownbrew/pkg/ownbrew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewInstall represents the install command
func NewInstall(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			cfg, err := pkgcmd.ReadConfig(l, cmd)
			if err != nil {
				return err
			}

			if err := viper.Unmarshal(&cfg); err != nil {
				return err
			}

			dry, err := cmd.Flags().GetBool("dry")
			if err != nil {
				return err
			}

			tags, err := cmd.Flags().GetStringSlice("tags")
			if err != nil {
				return err
			}

			brew, err := ownbrew.New(l,
				ownbrew.WithDry(dry),
				ownbrew.WithBinDir(cfg.BinDir),
				ownbrew.WithTapDir(cfg.TapDir),
				ownbrew.WithTempDir(cfg.TempDir),
				ownbrew.WithCellarDir(cfg.CellarDir),
				ownbrew.WithPackages(cfg.Packages...),
			)
			if err != nil {
				return err
			}
			return brew.Install(cmd.Context(), tags...)
		},
	}

	cmd.Flags().Bool("dry", false, "print out the taps that will be installed")
	cmd.Flags().StringSlice("tags", nil, "filter by tags (e.g. ci,-test)")

	root.AddCommand(cmd)
}
