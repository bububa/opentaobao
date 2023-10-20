package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialCurrencyGetAPIRequest 币种获取接口 API请求
// aliexpress.social.currency.get
//
// 获取目前AE社交支持的币种
type AliexpressSocialCurrencyGetAPIRequest struct {
	model.Params
}

// NewAliexpressSocialCurrencyGetRequest 初始化AliexpressSocialCurrencyGetAPIRequest对象
func NewAliexpressSocialCurrencyGetRequest() *AliexpressSocialCurrencyGetAPIRequest {
	return &AliexpressSocialCurrencyGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSocialCurrencyGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialCurrencyGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.currency.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialCurrencyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialCurrencyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAliexpressSocialCurrencyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSocialCurrencyGetRequest()
	},
}

// GetAliexpressSocialCurrencyGetRequest 从 sync.Pool 获取 AliexpressSocialCurrencyGetAPIRequest
func GetAliexpressSocialCurrencyGetAPIRequest() *AliexpressSocialCurrencyGetAPIRequest {
	return poolAliexpressSocialCurrencyGetAPIRequest.Get().(*AliexpressSocialCurrencyGetAPIRequest)
}

// ReleaseAliexpressSocialCurrencyGetAPIRequest 将 AliexpressSocialCurrencyGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSocialCurrencyGetAPIRequest(v *AliexpressSocialCurrencyGetAPIRequest) {
	v.Reset()
	poolAliexpressSocialCurrencyGetAPIRequest.Put(v)
}
