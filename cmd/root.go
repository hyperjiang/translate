package cmd

import (
	"os"

	"github.com/hyperjiang/translate/client"
	"github.com/hyperjiang/translate/translator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type options struct {
	inputFile  string
	outputFile string
	sourceLang string
	targetLang string
	listLang   bool
}

var opts options

func exit(err error) {
	log.Fatal().Err(err).Msg("")
}

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

func doTranslate(client client.Client, opts options) error {
	var jsonTranslator = translator.NewJSONTranslator(client)
	if err := jsonTranslator.ParseFile(opts.inputFile); err == nil {
		if err := jsonTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := jsonTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	var yamlTranslator = translator.NewYAMLTranslator(client)
	if err := yamlTranslator.ParseFile(opts.inputFile); err == nil {
		if err := yamlTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := yamlTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	return nil
}
