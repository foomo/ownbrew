package cmd

import (
	"encoding/json"
	"fmt"

	pkgcmd "github.com/foomo/ownbrew/pkg/cmd"
	"github.com/foomo/ownbrew/pkg/util"
	"github.com/spf13/cobra"
)

func NewConfig(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Print config",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			cfg, err := pkgcmd.ReadConfig(l, cmd)
			if err != nil {
				return err
			}

			out, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(util.Highlight(string(out)))
			return nil
		},
	}

	root.AddCommand(cmd)
}
