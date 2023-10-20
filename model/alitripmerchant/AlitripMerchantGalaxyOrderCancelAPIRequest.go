package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyordercancelAPIRequest 星河-取消预订 API请求
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
type AlitripmerchantgalaxyordercancelAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 订单编号
	_orderId string
	// 退款原因
	_reason string
}

// NewAlitripmerchantgalaxyordercancelRequest 初始化AlitripmerchantgalaxyordercancelAPIRequest对象
func NewAlitripmerchantgalaxyordercancelRequest() *AlitripmerchantgalaxyordercancelAPIRequest {
	return &AlitripmerchantgalaxyordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyordercancelAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyordercancelAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlitripmerchantgalaxyordercancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetReason is Reason Setter
// 退款原因
func (r *AlitripmerchantgalaxyordercancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlitripmerchantgalaxyordercancelAPIRequest) GetReason() string {
	return r._reason
}
