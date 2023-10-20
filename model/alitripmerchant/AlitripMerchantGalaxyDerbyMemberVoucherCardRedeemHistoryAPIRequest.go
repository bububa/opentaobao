package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest 查询会员兑换臻享卡历史记录 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户标识
	_token string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest.Put(v)
}
