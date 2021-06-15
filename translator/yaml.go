package translator

import (
	"io/ioutil"

	"github.com/hyperjiang/translate/client"
	"gopkg.in/yaml.v2"
)

type YAMLTranslator struct {
	Client   client.Client
	Original map[string]string
	Result   map[string]string
}

func NewYAMLTranslator(client client.Client) *YAMLTranslator {
	return &YAMLTranslator{Client: client}
}

func (trans *YAMLTranslator) ParseFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, &trans.Original)
}

func (trans *YAMLTranslator) Translate(sl, tl string) error {
	var err error
	trans.Result, err = translate(trans.Client, trans.Original, sl, tl)
	return err
}

func (trans *YAMLTranslator) SaveResult(file string) error {
	result, err := yaml.Marshal(trans.Result)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, result, 0644)
	return err
}
