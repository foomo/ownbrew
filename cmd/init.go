package cmd

import (
	"fmt"
	"os"
	"strings"

	pkgcmd "github.com/foomo/ownbrew/pkg/cmd"
	"github.com/foomo/ownbrew/pkg/config"
	"github.com/foomo/ownbrew/pkg/util"
	"github.com/spf13/cobra"
)

const initTemplate = `
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/ownbrew/refs/heads/main/ownbrew.schema.json
version: '%s'

binDir: "bin"
tapDir: ".ownbrew/tap"
tempDir: ".ownbrew/tmp"
cellarDir: ".ownbrew/bin"
packages:
  ## https://github.com/golangci/golangci-lint/releases
  - name: golangci-lint
    tags: [ci]
    tap: foomo/tap/golangci/golangci-lint
    version: 1.61.0
  ## https://github.com/go-courier/husky/releases
  - name: husky
    tap: foomo/tap/go-courier/husky
    version: 1.8.1
`

// NewInit represents the init command
func NewInit(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init ownbrew",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()
			filename := ".ownbrew.yaml"
			body := fmt.Sprintf(strings.Trim(initTemplate, "\n"), config.Version)

			if _, err := os.Stat(filename); err != nil && !os.IsNotExist(err) {
				return err
			} else if err == nil {
				l.Warn("ownbrew already initialized")
				fmt.Println(util.Highlight(body, "yaml"))
				return nil
			}

			return os.WriteFile(".ownbrew.yaml", []byte(body), 0600)
		},
	}

	root.AddCommand(cmd)
}
