package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
)

const (
	DefaultAliyunRegion = "cn-shenzhen"
	FormatTypeText      = "text"
	FormatTypeHTML      = "html"
)

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
		opts.RegionID = DefaultAliyunRegion
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

	request := newBatchTranslateRequest(string(jsonText), sl, tl, FormatTypeText)

	return c.BatchTranslate(request)
}

// api docs: https://next.api.aliyun.com/document/alimt/2018-10-12/GetBatchTranslate
func (c *AliyunClient) BatchTranslate(request *alimt.GetBatchTranslateRequest) (map[string]string, error) {
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

func newBatchTranslateRequest(st, sl, tl, ft string) *alimt.GetBatchTranslateRequest {
	request := alimt.CreateGetBatchTranslateRequest()
	request.Scheme = "https"
	request.Scene = "general"
	request.ApiType = "translate_standard"
	request.FormatType = ft
	request.SourceLanguage = sl
	request.TargetLanguage = tl
	request.SourceText = st

	return request
}
