package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatInfoAPIRequest 星河-获取微信用户的信息 API请求
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
type AlitripMerchantGalaxyWechatInfoAPIRequest struct {
	model.Params
	// 租户的id
	_tenantKey string
	// 微信小程序获取的code
	_code string
}

// NewAlitripMerchantGalaxyWechatInfoRequest 初始化AlitripMerchantGalaxyWechatInfoAPIRequest对象
func NewAlitripMerchantGalaxyWechatInfoRequest() *AlitripMerchantGalaxyWechatInfoAPIRequest {
	return &AlitripMerchantGalaxyWechatInfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatInfoAPIRequest) Reset() {
	r._tenantKey = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户的id
func (r *AlitripMerchantGalaxyWechatInfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCode is Code Setter
// 微信小程序获取的code
func (r *AlitripMerchantGalaxyWechatInfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetCode() string {
	return r._code
}

var poolAlitripMerchantGalaxyWechatInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatInfoRequest()
	},
}

// GetAlitripMerchantGalaxyWechatInfoRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatInfoAPIRequest
func GetAlitripMerchantGalaxyWechatInfoAPIRequest() *AlitripMerchantGalaxyWechatInfoAPIRequest {
	return poolAlitripMerchantGalaxyWechatInfoAPIRequest.Get().(*AlitripMerchantGalaxyWechatInfoAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatInfoAPIRequest 将 AlitripMerchantGalaxyWechatInfoAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatInfoAPIRequest(v *AlitripMerchantGalaxyWechatInfoAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatInfoAPIRequest.Put(v)
}
