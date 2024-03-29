package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest 查询用户拥有的臻享卡数量 API请求
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountRequest() *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.query.amount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest.Put(v)
}
