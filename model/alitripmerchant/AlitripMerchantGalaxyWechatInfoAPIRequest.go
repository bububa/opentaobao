package alitripmerchant

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 租户的id
func (r *AlitripMerchantGalaxyWechatInfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is Code Setter
// 微信小程序获取的code
func (r *AlitripMerchantGalaxyWechatInfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlitripMerchantGalaxyWechatInfoAPIRequest) GetCode() string {
	return r._code
}
