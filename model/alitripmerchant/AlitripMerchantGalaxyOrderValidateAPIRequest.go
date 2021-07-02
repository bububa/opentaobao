package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderValidateAPIRequest 星河-订单试单接口 API请求
// alitrip.merchant.galaxy.order.validate
//
// 根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
type AlitripMerchantGalaxyOrderValidateAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 试单参数
	_validateOrderParam *ValidateOrderParam
	// 用户标识
	_token string
}

// NewAlitripMerchantGalaxyOrderValidateRequest 初始化AlitripMerchantGalaxyOrderValidateAPIRequest对象
func NewAlitripMerchantGalaxyOrderValidateRequest() *AlitripMerchantGalaxyOrderValidateAPIRequest {
	return &AlitripMerchantGalaxyOrderValidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetValidateOrderParam is ValidateOrderParam Setter
// 试单参数
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetValidateOrderParam(_validateOrderParam *ValidateOrderParam) error {
	r._validateOrderParam = _validateOrderParam
	r.Set("validate_order_param", _validateOrderParam)
	return nil
}

// GetValidateOrderParam ValidateOrderParam Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetValidateOrderParam() *ValidateOrderParam {
	return r._validateOrderParam
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetToken() string {
	return r._token
}
