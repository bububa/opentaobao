package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest 德比线下权益券二维码查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
type AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户 token
	_token string
	// 会员权益ID
	_memberVoucherId string
}

// NewAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeRequest 初始化AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest对象
func NewAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeRequest() *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest {
	return &AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._memberVoucherId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户 token
func (r *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetMemberVoucherId is MemberVoucherId Setter
// 会员权益ID
func (r *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) SetMemberVoucherId(_memberVoucherId string) error {
	r._memberVoucherId = _memberVoucherId
	r.Set("member_voucher_id", _memberVoucherId)
	return nil
}

// GetMemberVoucherId MemberVoucherId Getter
func (r AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) GetMemberVoucherId() string {
	return r._memberVoucherId
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeRequest()
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeRequest 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest
func GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest() *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest 将 AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest(v *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest.Put(v)
}
