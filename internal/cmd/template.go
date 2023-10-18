package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/s-dwinter/yashiro/pkg/config"
	"github.com/s-dwinter/yashiro/pkg/engine"
	"github.com/spf13/cobra"
)

const example = `  # specify single file.
  ysr template example.yaml.tmpl

  # specify config file.
  ysr template -c config.yaml example.yaml.tmpl

  # specify multiple files using glob pattern.
  ysr template ./example/*.tmpl
`

func newTemplateCommand() *cobra.Command {
	var configFile string

	cmd := cobra.Command{
		Use:     "template <file>",
		Short:   "Generate a replaced text",
		Example: example,
		Args: func(_ *cobra.Command, args []string) error {
			return checkArgsLength(len(args), "template file")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			cfg, err := config.NewFromFile(ctx, configFile)
			if err != nil {
				return err
			}

			eng, err := engine.New(cfg)
			if err != nil {
				return err
			}

			b, err := readAllFiles(args[0])
			if err != nil {
				return err
			}

			return eng.Render(ctx, string(b), os.Stdout)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&configFile, "config", "c", config.DefaultConfigFilename, "specify config file.")

	return &cmd
}

func readAllFiles(pattern string) ([]byte, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("file not found: '%s'", pattern)
	}

	b := make([]byte, 0, 1024)
	for _, f := range files {
		c, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		b = append(b, c...)
	}

	return b, nil
}
