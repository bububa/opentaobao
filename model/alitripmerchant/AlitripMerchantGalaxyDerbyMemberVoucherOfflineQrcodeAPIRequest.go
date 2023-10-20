package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest 德比线下权益券二维码查询 API请求
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
type AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户 token
	_token string
	// 会员权益ID
	_memberVoucherId string
}

// NewAlitripmerchantgalaxyderbymembervoucherofflineqrcodeRequest 初始化AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherofflineqrcodeRequest() *AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户 token
func (r *AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetMemberVoucherId is MemberVoucherId Setter
// 会员权益ID
func (r *AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) SetMemberVoucherId(_memberVoucherId string) error {
	r._memberVoucherId = _memberVoucherId
	r.Set("member_voucher_id", _memberVoucherId)
	return nil
}

// GetMemberVoucherId MemberVoucherId Getter
func (r AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest) GetMemberVoucherId() string {
	return r._memberVoucherId
}
