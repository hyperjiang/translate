package translator

import (
	"errors"

	"github.com/hyperjiang/translate/client"
)

// Translator is the interface of translator
type Translator interface {
	ParseFile(file string) error
	Translate(sl, tl string) error
	SaveResult(file string) error
}

// maximum number of items per request
const MaxItemsPerRequest = 50

func translate(client client.Client, original map[string]string, sl, tl string) (map[string]string, error) {
	if client == nil {
		return nil, errors.New("translate client is nil")
	}

	if len(original) > MaxItemsPerRequest {
		firstPart := make(map[string]string)
		secondPart := make(map[string]string)
		var i int
		for k, v := range original {
			if i < MaxItemsPerRequest {
				firstPart[k] = v
			} else {
				secondPart[k] = v
			}
			i++
		}

		res1, err := client.Translate(firstPart, sl, tl)
		if err != nil {
			return nil, err
		}

		res2, err := translate(client, secondPart, sl, tl)
		if err != nil {
			return nil, err
		}

		// merge two parts of results
		for k, v := range res2 {
			res1[k] = v
		}

		return res1, nil
	}

	return client.Translate(original, sl, tl)
}
