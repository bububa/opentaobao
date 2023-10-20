package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatpaycallbackAPIRequest 微信支付回调 API请求
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
type AlitripmerchantgalaxywechatpaycallbackAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 回调参数
	_callBackJson string
}

// NewAlitripmerchantgalaxywechatpaycallbackRequest 初始化AlitripmerchantgalaxywechatpaycallbackAPIRequest对象
func NewAlitripmerchantgalaxywechatpaycallbackRequest() *AlitripmerchantgalaxywechatpaycallbackAPIRequest {
	return &AlitripmerchantgalaxywechatpaycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatpaycallbackAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.pay.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatpaycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatpaycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxywechatpaycallbackAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatpaycallbackAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCallBackJson is CallBackJson Setter
// 回调参数
func (r *AlitripmerchantgalaxywechatpaycallbackAPIRequest) SetCallBackJson(_callBackJson string) error {
	r._callBackJson = _callBackJson
	r.Set("call_back_json", _callBackJson)
	return nil
}

// GetCallBackJson CallBackJson Getter
func (r AlitripmerchantgalaxywechatpaycallbackAPIRequest) GetCallBackJson() string {
	return r._callBackJson
}
