package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialcountrygetAPIRequest 获取国家列表 API请求
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
type AliexpresssocialcountrygetAPIRequest struct {
	model.Params
	// 语言
	_language string
}

// NewAliexpresssocialcountrygetRequest 初始化AliexpresssocialcountrygetAPIRequest对象
func NewAliexpresssocialcountrygetRequest() *AliexpresssocialcountrygetAPIRequest {
	return &AliexpresssocialcountrygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialcountrygetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.country.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialcountrygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialcountrygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 语言
func (r *AliexpresssocialcountrygetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpresssocialcountrygetAPIRequest) GetLanguage() string {
	return r._language
}
