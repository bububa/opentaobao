package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest 订单详情查询接口 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 订单号
	_orderId string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest.Put(v)
}
