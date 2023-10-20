package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderCancelAPIRequest 星河-取消预订 API请求
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._orderId = ""
	r._reason = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetReason is Reason Setter
// 退款原因
func (r *AlitripMerchantGalaxyOrderCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlitripMerchantGalaxyOrderCancelAPIRequest) GetReason() string {
	return r._reason
}

var poolAlitripMerchantGalaxyOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyOrderCancelRequest()
	},
}

// GetAlitripMerchantGalaxyOrderCancelRequest 从 sync.Pool 获取 AlitripMerchantGalaxyOrderCancelAPIRequest
func GetAlitripMerchantGalaxyOrderCancelAPIRequest() *AlitripMerchantGalaxyOrderCancelAPIRequest {
	return poolAlitripMerchantGalaxyOrderCancelAPIRequest.Get().(*AlitripMerchantGalaxyOrderCancelAPIRequest)
}

// ReleaseAlitripMerchantGalaxyOrderCancelAPIRequest 将 AlitripMerchantGalaxyOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderCancelAPIRequest(v *AlitripMerchantGalaxyOrderCancelAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderCancelAPIRequest.Put(v)
}
