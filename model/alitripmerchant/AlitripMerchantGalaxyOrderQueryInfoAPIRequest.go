package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderqueryinfoAPIRequest 订单详情改版 API请求
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
type AlitripmerchantgalaxyorderqueryinfoAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户token
	_token string
	// 订单ID
	_orderId string
}

// NewAlitripmerchantgalaxyorderqueryinfoRequest 初始化AlitripmerchantgalaxyorderqueryinfoAPIRequest对象
func NewAlitripmerchantgalaxyorderqueryinfoRequest() *AlitripmerchantgalaxyorderqueryinfoAPIRequest {
	return &AlitripmerchantgalaxyorderqueryinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxyorderqueryinfoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyorderqueryinfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlitripmerchantgalaxyorderqueryinfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyorderqueryinfoAPIRequest) GetOrderId() string {
	return r._orderId
}
