package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest v5.0付费会员卡开发订单开票信息查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest struct {
	model.Params
	// 831232110dasdas
	_token string
	// 831232110
	_tenantKey string
	// 订单id
	_orderId string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) Reset() {
	r._token = ""
	r._tenantKey = ""
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 831232110dasdas
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetToken() string {
	return r._token
}

// SetTenantKey is TenantKey Setter
// 831232110
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest.Put(v)
}
