package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCardInfoAPIRequest 获取会员体系 API请求
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
type AlitripMerchantGalaxyCardInfoAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyCardInfoRequest 初始化AlitripMerchantGalaxyCardInfoAPIRequest对象
func NewAlitripMerchantGalaxyCardInfoRequest() *AlitripMerchantGalaxyCardInfoAPIRequest {
	return &AlitripMerchantGalaxyCardInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCardInfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.card.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCardInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCardInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyCardInfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCardInfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyCardInfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyCardInfoAPIRequest) GetToken() string {
	return r._token
}
