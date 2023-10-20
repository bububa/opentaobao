package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycardinfoAPIRequest 获取会员体系 API请求
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
type AlitripmerchantgalaxycardinfoAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripmerchantgalaxycardinfoRequest 初始化AlitripmerchantgalaxycardinfoAPIRequest对象
func NewAlitripmerchantgalaxycardinfoRequest() *AlitripmerchantgalaxycardinfoAPIRequest {
	return &AlitripmerchantgalaxycardinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycardinfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.card.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycardinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycardinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxycardinfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycardinfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxycardinfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxycardinfoAPIRequest) GetToken() string {
	return r._token
}
