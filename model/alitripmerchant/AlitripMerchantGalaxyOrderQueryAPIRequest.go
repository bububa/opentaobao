package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryAPIRequest 星河-单个订单详细信息查询 API请求
// alitrip.merchant.galaxy.order.query
//
// 为用户提供酒店订单的详细信息查询能力
type AlitripMerchantGalaxyOrderQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 订单号
	_orderId string
}

// NewAlitripMerchantGalaxyOrderQueryRequest 初始化AlitripMerchantGalaxyOrderQueryAPIRequest对象
func NewAlitripMerchantGalaxyOrderQueryRequest() *AlitripMerchantGalaxyOrderQueryAPIRequest {
	return &AlitripMerchantGalaxyOrderQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripMerchantGalaxyOrderQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyOrderQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlitripMerchantGalaxyOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyOrderQueryRequest()
	},
}

// GetAlitripMerchantGalaxyOrderQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyOrderQueryAPIRequest
func GetAlitripMerchantGalaxyOrderQueryAPIRequest() *AlitripMerchantGalaxyOrderQueryAPIRequest {
	return poolAlitripMerchantGalaxyOrderQueryAPIRequest.Get().(*AlitripMerchantGalaxyOrderQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyOrderQueryAPIRequest 将 AlitripMerchantGalaxyOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderQueryAPIRequest(v *AlitripMerchantGalaxyOrderQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderQueryAPIRequest.Put(v)
}
