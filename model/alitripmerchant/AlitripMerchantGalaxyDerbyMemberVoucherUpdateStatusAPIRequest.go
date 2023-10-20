package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest 前端订单支付成功回调-修改订单状态 API请求
// alitrip.merchant.galaxy.derby.member.voucher.update.status
//
// 前端订单支付成功回调-修改订单状态
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 订单ID
	_orderId string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusRequest() *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.update.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetToken() string {
	return r._token
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest.Put(v)
}
