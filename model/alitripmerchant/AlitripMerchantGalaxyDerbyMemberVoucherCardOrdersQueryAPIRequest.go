package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest 查询权益卡订单列表 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 状态
	_status string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.orders.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetToken() string {
	return r._token
}

// SetStatus is Status Setter
// 状态
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) GetStatus() string {
	return r._status
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrdersQueryAPIRequest.Put(v)
}
