package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialCountryGetAPIRequest 获取国家列表 API请求
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
type AliexpressSocialCountryGetAPIRequest struct {
	model.Params
	// 语言
	_language string
}

// NewAliexpressSocialCountryGetRequest 初始化AliexpressSocialCountryGetAPIRequest对象
func NewAliexpressSocialCountryGetRequest() *AliexpressSocialCountryGetAPIRequest {
	return &AliexpressSocialCountryGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSocialCountryGetAPIRequest) Reset() {
	r._language = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialCountryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.country.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialCountryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialCountryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 语言
func (r *AliexpressSocialCountryGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpressSocialCountryGetAPIRequest) GetLanguage() string {
	return r._language
}

var poolAliexpressSocialCountryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSocialCountryGetRequest()
	},
}

// GetAliexpressSocialCountryGetRequest 从 sync.Pool 获取 AliexpressSocialCountryGetAPIRequest
func GetAliexpressSocialCountryGetAPIRequest() *AliexpressSocialCountryGetAPIRequest {
	return poolAliexpressSocialCountryGetAPIRequest.Get().(*AliexpressSocialCountryGetAPIRequest)
}

// ReleaseAliexpressSocialCountryGetAPIRequest 将 AliexpressSocialCountryGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSocialCountryGetAPIRequest(v *AliexpressSocialCountryGetAPIRequest) {
	v.Reset()
	poolAliexpressSocialCountryGetAPIRequest.Put(v)
}
