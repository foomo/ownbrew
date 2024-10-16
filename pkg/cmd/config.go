package cmd

import (
	"bytes"
	"io"
	"log/slog"

	"github.com/foomo/ownbrew/pkg/config"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	viper.SetConfigType("yaml")
	filename := viper.GetString("config")

	if filename == "-" {
		return
	}
	if filename != "" {
		// Use config file from the flag.
		viper.SetConfigFile(filename)
		return
	}
	// Search config in home directory with name ".sesamy" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigName("ownbrew")
}

func ReadConfig(l *slog.Logger, cmd *cobra.Command) (*config.Config, error) {
	filename := viper.GetString("config")

	if filename == "-" {
		l.Debug("using config from stdin")
		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return nil, err
		}
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			return nil, err
		}
	} else {
		l.Debug("using config file", "filename", viper.ConfigFileUsed())
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}
	// l.Debug("config", l.ArgsFromMap(viper.AllSettings()))

	var cfg *config.Config
	if err := viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	}); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	if cfg.Version != config.Version {
		return nil, errors.New("missing or invalid config version: " + cfg.Version + " != '" + config.Version + "'")
	}

	return cfg, nil
}
