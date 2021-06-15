package translator

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hyperjiang/convert/client"
)

type JSONTranslator struct {
	Client   client.Client
	Original map[string]string
	Result   map[string]string
}

func NewJSONTranslator(client client.Client) *JSONTranslator {
	return &JSONTranslator{Client: client}
}

func (trans *JSONTranslator) ParseFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	trans.Original = make(map[string]string)

	return json.Unmarshal(content, &trans.Original)
}

func (trans *JSONTranslator) Translate(sl, tl string) error {
	var err error
	trans.Result, err = translate(trans.Client, trans.Original, sl, tl)

	return err
}

func (trans *JSONTranslator) SaveResult(file string) error {
	result, err := json.MarshalIndent(trans.Result, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, result, 0644)
}
