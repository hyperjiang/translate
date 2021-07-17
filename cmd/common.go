package cmd

import (
	"errors"

	"github.com/hyperjiang/translate/client"
	"github.com/hyperjiang/translate/translator"
	"github.com/rs/zerolog/log"
)

type options struct {
	inputFile  string
	outputFile string
	sourceLang string
	targetLang string
	listLang   bool
}

func exit(err error) {
	log.Fatal().Err(err).Msg("")
}

func exitf(format string, v ...interface{}) {
	log.Fatal().Msgf(format, v...)
}

func translate(client client.Client, opts options) error {
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

	var tsTranslator = translator.NewTsTranslator(client)
	if err := tsTranslator.ParseFile(opts.inputFile); err == nil {
		if err := tsTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := tsTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	var propTranslator = translator.NewPropTranslator(client)
	if err := propTranslator.ParseFile(opts.inputFile); err == nil {
		if err := propTranslator.Translate(opts.sourceLang, opts.targetLang); err != nil {
			return err
		}
		if err := propTranslator.SaveResult(opts.outputFile); err != nil {
			return err
		}
		return nil
	}

	return errors.New("unknown file format")
}
