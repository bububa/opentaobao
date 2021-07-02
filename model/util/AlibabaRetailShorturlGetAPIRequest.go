package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailShorturlGetAPIRequest 短链接获取 API请求
// alibaba.retail.shorturl.get
//
// 短链接获取
type AlibabaRetailShorturlGetAPIRequest struct {
	model.Params
	// 源url
	_sourceUrl string
	// 系统自动生成
	_options *ShortUrlOption
}

// NewAlibabaRetailShorturlGetRequest 初始化AlibabaRetailShorturlGetAPIRequest对象
func NewAlibabaRetailShorturlGetRequest() *AlibabaRetailShorturlGetAPIRequest {
	return &AlibabaRetailShorturlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailShorturlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.shorturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailShorturlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSourceUrl is SourceUrl Setter
// 源url
func (r *AlibabaRetailShorturlGetAPIRequest) SetSourceUrl(_sourceUrl string) error {
	r._sourceUrl = _sourceUrl
	r.Set("source_url", _sourceUrl)
	return nil
}

// GetSourceUrl SourceUrl Getter
func (r AlibabaRetailShorturlGetAPIRequest) GetSourceUrl() string {
	return r._sourceUrl
}

// SetOptions is Options Setter
// 系统自动生成
func (r *AlibabaRetailShorturlGetAPIRequest) SetOptions(_options *ShortUrlOption) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r AlibabaRetailShorturlGetAPIRequest) GetOptions() *ShortUrlOption {
	return r._options
}
