package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialCurrencyGetAPIRequest
币种获取接口 API请求
aliexpress.social.currency.get

获取目前AE社交支持的币种 */
type AliexpressSocialCurrencyGetAPIRequest struct {
	model.Params
}

// NewAliexpressSocialCurrencyGetRequest 初始化AliexpressSocialCurrencyGetAPIRequest对象
func NewAliexpressSocialCurrencyGetRequest() *AliexpressSocialCurrencyGetAPIRequest {
	return &AliexpressSocialCurrencyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialCurrencyGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.currency.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialCurrencyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
