package translator

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/hyperjiang/translate/client"
)

type TsTranslator struct {
	Client   client.Client
	Original map[string]string
	Result   map[string]string
}

func NewTsTranslator(client client.Client) *TsTranslator {
	return &TsTranslator{Client: client}
}

func (trans *TsTranslator) ParseFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	str := strings.Replace(string(content), "export default", "", -1)
	str = strings.Replace(str, "'", "\"", -1)

	exp := `,\s+\};?`
	reg := regexp.MustCompile(exp)
	m := reg.FindStringSubmatch(str)
	if len(m) > 0 {
		str = strings.Replace(str, m[0], "}", -1)
	}

	trans.Original = make(map[string]string)

	return json.Unmarshal([]byte(str), &trans.Original)
}

func (trans *TsTranslator) Translate(sl, tl string) error {
	var err error
	trans.Result, err = translate(trans.Client, trans.Original, sl, tl)

	return err
}

func (trans *TsTranslator) SaveResult(file string) error {
	result, err := json.MarshalIndent(trans.Result, "", "  ")
	if err != nil {
		return err
	}

	s := "export default " + string(result) + ";"
	s = strings.Replace(s, "\"", "'", -1)
	res := []byte(s)

	return ioutil.WriteFile(file, res, 0644)
}
