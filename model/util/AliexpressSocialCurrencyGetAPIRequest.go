package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialcurrencygetAPIRequest 币种获取接口 API请求
// aliexpress.social.currency.get
//
// 获取目前AE社交支持的币种
type AliexpresssocialcurrencygetAPIRequest struct {
	model.Params
}

// NewAliexpresssocialcurrencygetRequest 初始化AliexpresssocialcurrencygetAPIRequest对象
func NewAliexpresssocialcurrencygetRequest() *AliexpresssocialcurrencygetAPIRequest {
	return &AliexpresssocialcurrencygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialcurrencygetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.currency.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialcurrencygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialcurrencygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
