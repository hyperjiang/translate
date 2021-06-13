package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
)

const defaultAliyunRegion = "cn-shenzhen"

type AliyunClient struct {
	Client *alimt.Client
}

type AliyunOptions struct {
	RegionID     string
	AccessKeyID  string
	AccessSecret string
}

func NewAliyunClient(opts AliyunOptions) (*AliyunClient, error) {
	if opts.RegionID == "" {
		opts.RegionID = defaultAliyunRegion
	}

	client, err := alimt.NewClientWithAccessKey(opts.RegionID, opts.AccessKeyID, opts.AccessSecret)

	if err != nil {
		return nil, err
	}

	return &AliyunClient{Client: client}, nil
}

func (c *AliyunClient) Translate(original map[string]string, sl, tl string) (map[string]string, error) {
	jsonText, err := json.Marshal(original)
	if err != nil {
		return nil, err
	}

	request := alimt.CreateGetBatchTranslateRequest()
	request.Scheme = "https"
	request.FormatType = "text"
	request.Scene = "general"
	request.ApiType = "translate_standard"
	request.SourceLanguage = sl
	request.TargetLanguage = tl
	request.SourceText = string(jsonText)

	response, err := c.Client.GetBatchTranslate(request)
	if err != nil {
		return nil, err
	}

	if response.Code != 200 {
		return nil, errors.New(response.Message)
	}

	result := make(map[string]string)
	for _, m := range response.TranslatedList {
		k := fmt.Sprintf("%s", m["index"])
		v := fmt.Sprintf("%s", m["translated"])
		result[k] = v
	}

	return result, nil
}

func (c *AliyunClient) ListSupportedLanguages() {
	fmt.Println("See https://help.aliyun.com/document_detail/158269.html")
}
