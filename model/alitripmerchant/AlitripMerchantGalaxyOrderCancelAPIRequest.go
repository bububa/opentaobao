package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyOrderCancelAPIRequest
星河-取消预订 API请求
alitrip.merchant.galaxy.order.cancel

雅高酒店用户使用该接口，取消酒店预订 */
type AlitripMerchantGalaxyOrderCancelAPIRequest struct {
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

// NewAlitripMerchantGalaxyOrderCancelRequest 初始化AlitripMerchantGalaxyOrderCancelAPIRequest对象
func NewAlitripMerchantGalaxyOrderCancelRequest() *AlitripMerchantGalaxyOrderCancelAPIRequest {
	return &AlitripMerchantGalaxyOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetToken() string {
	return r._token
}

// Set is OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Reason Setter
// 退款原因
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetReason() string {
	return r._reason
}
