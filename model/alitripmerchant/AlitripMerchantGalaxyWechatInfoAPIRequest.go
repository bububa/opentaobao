package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatinfoAPIRequest 星河-获取微信用户的信息 API请求
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
type AlitripmerchantgalaxywechatinfoAPIRequest struct {
	model.Params
	// 租户的id
	_tenantKey string
	// 微信小程序获取的code
	_code string
}

// NewAlitripmerchantgalaxywechatinfoRequest 初始化AlitripmerchantgalaxywechatinfoAPIRequest对象
func NewAlitripmerchantgalaxywechatinfoRequest() *AlitripmerchantgalaxywechatinfoAPIRequest {
	return &AlitripmerchantgalaxywechatinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatinfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户的id
func (r *AlitripmerchantgalaxywechatinfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatinfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCode is Code Setter
// 微信小程序获取的code
func (r *AlitripmerchantgalaxywechatinfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripmerchantgalaxywechatinfoAPIRequest) GetCode() string {
	return r._code
}
