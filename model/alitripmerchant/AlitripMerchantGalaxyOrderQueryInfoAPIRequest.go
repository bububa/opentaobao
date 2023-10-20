package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryInfoAPIRequest 订单详情改版 API请求
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
type AlitripMerchantGalaxyOrderQueryInfoAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户token
	_token string
	// 订单ID
	_orderId string
}

// NewAlitripMerchantGalaxyOrderQueryInfoRequest 初始化AlitripMerchantGalaxyOrderQueryInfoAPIRequest对象
func NewAlitripMerchantGalaxyOrderQueryInfoRequest() *AlitripMerchantGalaxyOrderQueryInfoAPIRequest {
	return &AlitripMerchantGalaxyOrderQueryInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyOrderQueryInfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyOrderQueryInfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlitripMerchantGalaxyOrderQueryInfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyOrderQueryInfoAPIRequest) GetOrderId() string {
	return r._orderId
}
