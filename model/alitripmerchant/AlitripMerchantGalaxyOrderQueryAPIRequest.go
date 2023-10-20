package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderqueryAPIRequest 星河-单个订单详细信息查询 API请求
// alitrip.merchant.galaxy.order.query
//
// 为用户提供酒店订单的详细信息查询能力
type AlitripmerchantgalaxyorderqueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 订单号
	_orderId string
}

// NewAlitripmerchantgalaxyorderqueryRequest 初始化AlitripmerchantgalaxyorderqueryAPIRequest对象
func NewAlitripmerchantgalaxyorderqueryRequest() *AlitripmerchantgalaxyorderqueryAPIRequest {
	return &AlitripmerchantgalaxyorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyorderqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyorderqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripmerchantgalaxyorderqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyorderqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
