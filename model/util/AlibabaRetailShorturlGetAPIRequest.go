package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailShorturlGetAPIRequest) Reset() {
	r._sourceUrl = ""
	r._options = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailShorturlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.shorturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailShorturlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailShorturlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaRetailShorturlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailShorturlGetRequest()
	},
}

// GetAlibabaRetailShorturlGetRequest 从 sync.Pool 获取 AlibabaRetailShorturlGetAPIRequest
func GetAlibabaRetailShorturlGetAPIRequest() *AlibabaRetailShorturlGetAPIRequest {
	return poolAlibabaRetailShorturlGetAPIRequest.Get().(*AlibabaRetailShorturlGetAPIRequest)
}

// ReleaseAlibabaRetailShorturlGetAPIRequest 将 AlibabaRetailShorturlGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailShorturlGetAPIRequest(v *AlibabaRetailShorturlGetAPIRequest) {
	v.Reset()
	poolAlibabaRetailShorturlGetAPIRequest.Put(v)
}
