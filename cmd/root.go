package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var opts options

var rootCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translator",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exit(err)
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	viper.AutomaticEnv()

	rootCmd.AddCommand(aliyunCmd)

	rootCmd.PersistentFlags().StringVarP(&opts.inputFile, "input", "i", "", "the input file to be translated, must provide")
	rootCmd.PersistentFlags().StringVarP(&opts.outputFile, "output", "o", "/dev/stdout", "the output path to save translated file")
	rootCmd.PersistentFlags().StringVarP(&opts.sourceLang, "source", "s", "en", "source language")
	rootCmd.PersistentFlags().StringVarP(&opts.targetLang, "targe", "t", "en", "target language")
	rootCmd.PersistentFlags().BoolVarP(&opts.listLang, "languages", "l", false, "list available languages")
}
