package util

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialCountryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.country.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialCountryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
