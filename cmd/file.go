package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File format converter",
	Run: func(cmd *cobra.Command, args []string) {
		if opts.inputFile != "" && opts.outputFile != "" {
			content, err := ioutil.ReadFile(opts.inputFile)
			if err != nil {
				exit(err)
			}

			var data interface{}
			ext := filepath.Ext(opts.inputFile)
			switch ext {
			case ".json":
				json.Unmarshal(content, &data)
			case ".yaml", ".yml":
				yaml.Unmarshal(content, &data)
			default:
				exitf("unknown input file extension: %s", ext)
			}

			var result []byte
			ext2 := filepath.Ext(opts.outputFile)
			switch ext2 {
			case ".json":
				if result, err = json.MarshalIndent(data, "", "    "); err != nil {
					exit(err)
				}
			case ".yaml", ".yml":
				if result, err = yaml.Marshal(data); err != nil {
					exit(err)
				}
			default:
				exitf("unknown output file extension: %s", ext2)
			}

			if err := ioutil.WriteFile(opts.outputFile, result, 0644); err != nil {
				exit(err)
			}

			log.Info().Msgf("convert successfully and save into %s", opts.outputFile)

		} else {
			cmd.Usage()
		}
	},
}

func init() {
	fileCmd.PersistentFlags().StringVarP(&opts.inputFile, "input", "i", "", "the input file to be converted, must provide")
	fileCmd.PersistentFlags().StringVarP(&opts.outputFile, "output", "o", "", "the output path to save result, must provide")
}
