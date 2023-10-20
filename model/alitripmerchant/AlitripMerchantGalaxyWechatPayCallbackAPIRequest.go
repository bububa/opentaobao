package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatPayCallbackAPIRequest 微信支付回调 API请求
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
type AlitripMerchantGalaxyWechatPayCallbackAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 回调参数
	_callBackJson string
}

// NewAlitripMerchantGalaxyWechatPayCallbackRequest 初始化AlitripMerchantGalaxyWechatPayCallbackAPIRequest对象
func NewAlitripMerchantGalaxyWechatPayCallbackRequest() *AlitripMerchantGalaxyWechatPayCallbackAPIRequest {
	return &AlitripMerchantGalaxyWechatPayCallbackAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatPayCallbackAPIRequest) Reset() {
	r._tenantKey = ""
	r._callBackJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatPayCallbackAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.pay.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatPayCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatPayCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripMerchantGalaxyWechatPayCallbackAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatPayCallbackAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCallBackJson is CallBackJson Setter
// 回调参数
func (r *AlitripMerchantGalaxyWechatPayCallbackAPIRequest) SetCallBackJson(_callBackJson string) error {
	r._callBackJson = _callBackJson
	r.Set("call_back_json", _callBackJson)
	return nil
}

// GetCallBackJson CallBackJson Getter
func (r AlitripMerchantGalaxyWechatPayCallbackAPIRequest) GetCallBackJson() string {
	return r._callBackJson
}

var poolAlitripMerchantGalaxyWechatPayCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatPayCallbackRequest()
	},
}

// GetAlitripMerchantGalaxyWechatPayCallbackRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatPayCallbackAPIRequest
func GetAlitripMerchantGalaxyWechatPayCallbackAPIRequest() *AlitripMerchantGalaxyWechatPayCallbackAPIRequest {
	return poolAlitripMerchantGalaxyWechatPayCallbackAPIRequest.Get().(*AlitripMerchantGalaxyWechatPayCallbackAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatPayCallbackAPIRequest 将 AlitripMerchantGalaxyWechatPayCallbackAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatPayCallbackAPIRequest(v *AlitripMerchantGalaxyWechatPayCallbackAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatPayCallbackAPIRequest.Put(v)
}
