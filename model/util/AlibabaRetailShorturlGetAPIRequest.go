package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailshorturlgetAPIRequest 短链接获取 API请求
// alibaba.retail.shorturl.get
//
// 短链接获取
type AlibabaretailshorturlgetAPIRequest struct {
	model.Params
	// 源url
	_sourceUrl string
	// 系统自动生成
	_options *ShortUrlOption
}

// NewAlibabaretailshorturlgetRequest 初始化AlibabaretailshorturlgetAPIRequest对象
func NewAlibabaretailshorturlgetRequest() *AlibabaretailshorturlgetAPIRequest {
	return &AlibabaretailshorturlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailshorturlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.shorturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailshorturlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailshorturlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceUrl is SourceUrl Setter
// 源url
func (r *AlibabaretailshorturlgetAPIRequest) SetSourceUrl(_sourceUrl string) error {
	r._sourceUrl = _sourceUrl
	r.Set("source_url", _sourceUrl)
	return nil
}

// GetSourceUrl SourceUrl Getter
func (r AlibabaretailshorturlgetAPIRequest) GetSourceUrl() string {
	return r._sourceUrl
}

// SetOptions is Options Setter
// 系统自动生成
func (r *AlibabaretailshorturlgetAPIRequest) SetOptions(_options *ShortUrlOption) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r AlibabaretailshorturlgetAPIRequest) GetOptions() *ShortUrlOption {
	return r._options
}
