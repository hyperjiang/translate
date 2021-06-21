package translator

import (
	"io/ioutil"

	"github.com/hyperjiang/translate/client"
	"github.com/magiconair/properties"
)

type PropTranslator struct {
	Client   client.Client
	Original map[string]string
	Result   map[string]string
}

func NewPropTranslator(client client.Client) *PropTranslator {
	return &PropTranslator{Client: client}
}

func (trans *PropTranslator) ParseFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	p, err := properties.LoadString(string(content))
	if err != nil {
		return nil
	}

	trans.Original = p.Map()

	return nil
}

func (trans *PropTranslator) Translate(sl, tl string) error {
	var err error
	trans.Result, err = translate(trans.Client, trans.Original, sl, tl)

	return err
}

func (trans *PropTranslator) SaveResult(file string) error {
	result := BuildProperties(trans.Result)

	return ioutil.WriteFile(file, result, 0644)
}
